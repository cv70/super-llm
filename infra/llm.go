package infra

import (
	"context"
	"super-llm/config"
	"super-llm/pkg/sdk"

	"google.golang.org/adk/model"
)

func NewLLM(ctx context.Context, c *config.LLMConfig) (model.LLM, error) {
	model, err := sdk.NewModel(ctx, c.Model, &sdk.ClientConfig{
		BaseURL: c.BaseURL,
		APIKey: c.APIKey,
	})
	return model, err
}
