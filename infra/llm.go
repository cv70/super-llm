package infra

import (
	"context"
	"super-llm/config"

	"github.com/cv70/pkgo/llm"
)

func NewLLM(ctx context.Context, c *config.LLMConfig) (*llm.OpenAIModel, error) {
	model, err := llm.NewModel(ctx, c.Model, &llm.ClientConfig{
		BaseURL: c.BaseURL,
		APIKey:  c.APIKey,
	})
	return model, err
}
