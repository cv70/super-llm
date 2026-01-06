package committee

import (
	"context"
	"super-llm/config"
	"super-llm/infra"

	"github.com/pkg/errors"
	"google.golang.org/adk/model"
)

type CommitteeDomain struct {
	Members map[string]model.LLM
}

func BuildCommitteeDomain(ctx context.Context, cfg *config.Config) (*CommitteeDomain, error) {
	if cfg == nil || len(cfg.LLMs) == 0 {
		return nil, errors.New("invalid configuration")
	}

	domain := &CommitteeDomain{
		Members:    map[string]model.LLM{},
	}

	// Initialize members
	for _, llmCfg := range cfg.LLMs {
		model, err := infra.NewLLM(ctx, llmCfg)
		if err != nil {
			return nil, errors.Errorf("failed to create LLM for %s: %v", llmCfg.Model, err)
		}
		
		domain.Members[model.Name()] = model
	}
	return domain, nil
}