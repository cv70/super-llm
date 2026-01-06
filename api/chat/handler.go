package chat

import (
	"context"
	"log/slog"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"

	"super-llm/domain/committee"
	"super-llm/pkg/sdk"
)

// Handler holds dependencies for the chat handler
type Handler struct {
	committee  *committee.CommitteeDomain
}

// NewHandler creates a new chat handler
func NewHandler(committee *committee.CommitteeDomain) *Handler {
	return &Handler{
		committee:  committee,
	}
}

// ChatCompletions handles the /chat/completions endpoint
func (h *Handler) ChatCompletions(c *gin.Context) {
	// Parse request body
	var req sdk.ChatCompletionRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	// Parse headers
	membersHeader := c.GetHeader("X-Members")
	viewsHeader := c.GetHeader("X-Views")

	if membersHeader == "" {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal request header"})
		return
	}

	// Process models from header
	members := parseModels(membersHeader)

	// Process stage from header
	// opinion, review
	var opinion, review bool
	if viewsHeader != "" {
		opinion, review = parseViews(viewsHeader)
	}

	// Log the request
	slog.Info(
		"Received chat completions request",
		slog.Any("members", membersHeader),
		slog.Any("views", viewsHeader),
		slog.Any("messages_count", len(req.Messages)),
	)

	// Process the request using committee and LLM service
	response, err := h.processRequest(c, &req, members, opinion, review)
	if err != nil {
		slog.Error("Failed to process chat completions", slog.Any("err", err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	// Return response
	c.JSON(http.StatusOK, response)
}

// processRequest processes the chat completion request
func (h *Handler) processRequest(ctx context.Context, req *sdk.ChatCompletionRequest, members []string, opinion, review bool) (*sdk.ChatCompletionResponse, error) {
	// For simplicity, we'll use the RunCommitteeProcess method directly
	// Since our interface requires a single question, we'll just use the first user message
	result, err := h.committee.RunCommitteeProcess(ctx, req, members, opinion, review)
	if err != nil {
		return nil, err
	}

	// Build response
	response := &sdk.ChatCompletionResponse{
		ID:      "chatcmpl-" + time.Now().Format("20060102150405"),
		Object:  "chat.completion",
		Created: time.Now().Unix(),
		Model:   req.Model,
		Choices: []sdk.ChatChoice{{
			Index: 0,
			Message: &sdk.ChatMessage{
				Role:    "assistant",
				Content: result,
			},
			FinishReason: "stop",
		}},
		Usage: &sdk.ChatUsage{
			PromptTokens:     0,
			CompletionTokens: 0,
			TotalTokens:      0,
		},
	}

	return response, nil
}

// parseModels parses comma-separated model names from header
func parseModels(header string) []string {
	// Simple split by comma, trim spaces
	models := []string{}
	for _, model := range strings.Split(header, ",") {
		model = strings.TrimSpace(model)
		if model != "" {
			models = append(models, model)
		}
	}
	return models
}

// parseViews parses view from header
func parseViews(header string) (bool, bool) {
	parts := strings.Split(header, ",")
	opinion := false
	review := false
	for _, part := range parts {
		switch strings.ToLower(strings.TrimSpace(part)) {
		case "opinion":
			opinion = true
		case "review":
			review = true
		}
	}
	return opinion, review
}