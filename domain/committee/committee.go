package committee

import (
	"context"
	"fmt"
	"iter"
	"log/slog"
	"maps"
	"strings"
	"super-llm/pkg/sdk"
	"sync"

	"github.com/pkg/errors"
	"google.golang.org/adk/model"
	"google.golang.org/genai"
)

// GetMembers returns the list of committee members
func (d *CommitteeDomain) GetMembers() iter.Seq[model.LLM] {
	return maps.Values(d.Members)
}

// Phase1InitialOpinions collects initial opinions from all LLMs
func (d *CommitteeDomain) Phase1InitialOpinions(c *CommitteeContext) error {
	results := make(map[string]string)
	var wg sync.WaitGroup
	var mu sync.Mutex
	
	// Create a channel to collect results
	resultChan := make(chan struct {
		name  string
		reply string
		err   error
	}, len(d.Members))

	// Send question to all LLMs concurrently
	for member := range d.GetMembers() {
		wg.Add(1)
		go func(member model.LLM) {
			defer wg.Done()
			
			// Create content for the LLM
			content := genai.NewContentFromText(c.MessageSummary, genai.RoleUser)
			
			// Create request
			req := &model.LLMRequest{
				Contents: []*genai.Content{content},
			}
			
			// Generate response
			seq := member.GenerateContent(c, req, false)
			var response *model.LLMResponse
			for resp, err := range seq {
				if err != nil {
					resultChan <- struct {
						name  string
						reply string
						err   error
					}{member.Name(), "", err}
					return
				}
				response = resp
				break
			}
			
			if response == nil {
				resultChan <- struct {
					name  string
					reply string
					err   error
				}{member.Name(), "", fmt.Errorf("no response from %v", member.Name())}
				return
			}
			
			// Extract text from response
			var replyText string
			if response.Content != nil {
				for _, part := range response.Content.Parts {
					if part.Text != "" {
						replyText += part.Text
					}
				}
			}
			
			resultChan <- struct {
				name  string
				reply string
				err   error
			}{member.Name(), replyText, nil}
		}(member)
	}

	// Close result channel when all goroutines are done
	go func() {
		wg.Wait()
		close(resultChan)
	}()

	// Collect results
	for result := range resultChan {
		mu.Lock()
		results[result.name] = result.reply
		mu.Unlock()
		
		if result.err != nil {
			slog.Error("getting opinion", slog.Any("name", result.name), slog.Any("err", result.err))
		}
	}

	c.Opinions = results
	return nil
}

// Phase2Review evaluates and ranks all responses anonymously
func (d *CommitteeDomain) Phase2Review(c *CommitteeContext) error {
	reviews := make(map[string][]string)
	var wg sync.WaitGroup
	var mu sync.Mutex
	
	// Create a channel to collect reviews
	reviewChan := make(chan struct {
		name   string
		review []string
		err    error
	}, len(d.Members))

	// Each LLM reviews all other LLMs' responses anonymously
	for member := range d.GetMembers() {
		wg.Add(1)
		go func(member model.LLM) {
			defer wg.Done()
			
			// Prepare review prompt
			var promptBuilder strings.Builder
			promptBuilder.WriteString("请对以下内容的回复进行匿名评审和排名：\n\n")
			promptBuilder.WriteString(c.MessageSummary)
			promptBuilder.WriteString("\n\n请对以下回复进行评分和排名（从高到低）：\n")
			
			// Add all opinions
			for name, opinion := range c.Opinions {
				promptBuilder.WriteString(name)
				promptBuilder.WriteString(": ")
				promptBuilder.WriteString(opinion)
				promptBuilder.WriteString("\n\n")
			}
			
			promptBuilder.WriteString("请按以下格式回答：\n")
			promptBuilder.WriteString("1. 评分最高的回复：[模型名称]\n")
			promptBuilder.WriteString("2. 评分第二高的回复：[模型名称]\n")
			promptBuilder.WriteString("3. 评分第三高的回复：[模型名称]\n")
			promptBuilder.WriteString("4. 详细评价：[简要说明]\n")
			
			// Create content for the LLM
			content := genai.NewContentFromText(promptBuilder.String(), genai.RoleUser)
			
			// Create request
			req := &model.LLMRequest{
				Contents: []*genai.Content{content},
			}
			
			// Generate response
			seq := member.GenerateContent(c, req, false)
			var response *model.LLMResponse
			for resp, err := range seq {
				if err != nil {
					reviewChan <- struct {
						name   string
						review []string
						err    error
					}{member.Name(), nil, err}
					return
				}
				response = resp
				break
			}
			
			if response == nil {
				reviewChan <- struct {
					name   string
					review []string
					err    error
				}{member.Name(), nil, errors.Errorf("no response from %v", member.Name())}
				return
			}
			
			// Extract text from response
			var reviewText string
			if response.Content != nil {
				for _, part := range response.Content.Parts {
					if part.Text != "" {
						reviewText += part.Text
					}
				}
			}
			
			// Parse review response
			reviewLines := strings.Split(reviewText, "\n")
			review := make([]string, len(reviewLines))
			for i, line := range reviewLines {
				review[i] = strings.TrimSpace(line)
			}
			
			reviewChan <- struct {
				name   string
				review []string
				err    error
			}{member.Name(), review, nil}
		}(member)
	}

	// Close review channel when all goroutines are done
	go func() {
		wg.Wait()
		close(reviewChan)
	}()

	// Collect reviews
	for result := range reviewChan {
		mu.Lock()
		reviews[result.name] = result.review
		mu.Unlock()
		
		if result.err != nil {
			slog.Error("getting review", slog.Any("name", result.name), slog.Any("err", result.err))
		}
	}

	c.Reviews = reviews
	return nil
}

// Phase3FinalAnswer generates the final answer using the leader model
func (d *CommitteeDomain) Phase3FinalAnswer(c *CommitteeContext) (string, error) {
	// Prepare final answer prompt
	var promptBuilder strings.Builder
	promptBuilder.WriteString("请基于以下信息生成最终回答：\n\n")
	promptBuilder.WriteString("需求：")
	promptBuilder.WriteString(c.MessageSummary)
	promptBuilder.WriteString("\n\n")
	
	// // Get the question from the first user message
	// if len(c.Messages) > 0 {
	// 	for _, msg := range c.Messages {
	// 		if msg.Role == "user" {
	// 			if contentStr, ok := msg.Content.(string); ok {
	// 				promptBuilder.WriteString(fmt.Sprintf("问题：%s\n\n", contentStr))
	// 				break
	// 			}
	// 		}
	// 	}
	// }
	
	promptBuilder.WriteString("各模型的初始回复：\n")
	for name, opinion := range c.Opinions {
		promptBuilder.WriteString(fmt.Sprintf("%s: %s\n\n", name, opinion))
	}
	
	promptBuilder.WriteString("各模型的评审意见：\n")
	for name, review := range c.Reviews {
		promptBuilder.WriteString(fmt.Sprintf("%s 的评审：\n", name))
		for _, line := range review {
			promptBuilder.WriteString(fmt.Sprintf("  %s\n", line))
		}
		promptBuilder.WriteString("\n")
	}
	
	promptBuilder.WriteString("请综合所有回复和评审意见，给出一个高质量、准确且全面的最终回答。")

	// Create content for the leader model
	content := genai.NewContentFromText(promptBuilder.String(), "user")
	
	// Create request
	req := &model.LLMRequest{
		Contents: []*genai.Content{content},
	}
	
	// Generate response
	seq := c.Leader.GenerateContent(c, req, false)
	var response *model.LLMResponse
	for resp, err := range seq {
		if err != nil {
			return "", err
		}
		response = resp
		break
	}
	
	if response == nil {
		return "", errors.Errorf("no response from leader model")
	}
	
	// Extract text from response
	var answerText string
	if response.Content != nil {
		for _, part := range response.Content.Parts {
			if part.Text != "" {
				answerText += part.Text
			}
		}
	}
	
	return answerText, nil
}

// GenerateConversationSummary generates a summary of the conversation
func (d *CommitteeDomain) GenerateConversationSummary(c *CommitteeContext) error {
	if len(c.Messages) == 0 {
		return nil
	}

	// Prepare summary prompt
	var promptBuilder strings.Builder
	promptBuilder.WriteString("请总结以下对话内容，提取关键信息和要点：\n\n")

	// Add all messages to the summary prompt
	for _, message := range c.Messages {
		if message.Role == "user" {
			promptBuilder.WriteString("用户问题：")
		} else if message.Role == "assistant" {
			promptBuilder.WriteString("助手回答：")
		} else {
			promptBuilder.WriteString("系统信息：")
		}
		if contentStr, ok := message.Content.(string); ok {
			promptBuilder.WriteString(contentStr)
		} else if contentArr, ok := message.Content.([]any); ok {
			// Handle array content
			for _, item := range contentArr {
				if str, ok := item.(string); ok {
					promptBuilder.WriteString(str)
				}
			}
		}
		promptBuilder.WriteString("\n\n")
	}
	
	promptBuilder.WriteString("请用简洁明了的语言总结以上对话的主要内容和关键点。")

	// Create content for the leader model
	content := genai.NewContentFromText(promptBuilder.String(), "user")
	
	// Create request
	req := &model.LLMRequest{
		Contents: []*genai.Content{content},
	}
	
	// Generate response
	seq := c.Leader.GenerateContent(c, req, false)
	var response *model.LLMResponse
	for resp, err := range seq {
		if err != nil {
			return err
		}
		response = resp
		break
	}
	
	if response == nil {
		return errors.New("no response from leader model for summary")
	}
	
	// Extract text from response
	var summaryText string
	if response.Content != nil {
		for _, part := range response.Content.Parts {
			if part.Text != "" {
				summaryText += part.Text
			}
		}
	}
	
	c.MessageSummary = summaryText
	return nil
}

// RunCommitteeProcess executes the complete committee process
func (d *CommitteeDomain) RunCommitteeProcess(ctx context.Context, req *sdk.ChatCompletionRequest, members []string, opinion, review bool) (string, error) {
	c, err := d.BuildCommitteeContext(ctx, req, members, opinion, review)
	if err != nil {
		return "", errors.Wrap(err, "build committee context")
	}
	slog.Info("LLM 委员会开始处理问题...", slog.Any("message", c.Messages))
	
	// Generate summary before phase 1
	slog.Info("生成对话摘要...")
	err = d.GenerateConversationSummary(c)
	if err != nil {
		slog.Error("生成摘要失败", slog.Any("err", err))
		return "", errors.Wrap(err, "generate summary")
	}
	
	// Phase 1: Initial Opinions
	slog.Info("第一阶段：收集初步意见...")
	err = d.Phase1InitialOpinions(c)
	if err != nil {
		return "", errors.Wrap(err, "phase 1")
	}

	// Phase 2: Review
	slog.Info("\n第二阶段：交叉评审...")
	err = d.Phase2Review(c)
	if err != nil {
		return "", errors.Wrap(err, "phase 2")
	}
	
	// Phase 3: Final Answer
	slog.Info("第三阶段：生成最终答案...")
	finalAnswer, err := d.Phase3FinalAnswer(c)
	if err != nil {
		return "", errors.Wrap(err, "phase 3")
	}
	
	slog.Info("最终答案：")
	slog.Info(finalAnswer)
	
	return finalAnswer, nil
}