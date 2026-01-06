package committee

import (
	"context"
	"super-llm/pkg/gslice"
	"super-llm/pkg/sdk"

	"github.com/pkg/errors"
	"google.golang.org/adk/model"
)

type CommitteeContext struct {
	context.Context
	Messages []*sdk.ChatMessage
	Leader model.LLM
	Members map[string]model.LLM

	Opinions map[string]string
	Reviews map[string][]string
	MessageSummary string

	OutputOpinion bool
	OutputReview bool
}

func (d *CommitteeDomain) BuildCommitteeContext(ctx context.Context, req *sdk.ChatCompletionRequest, members []string, opinion, review bool) (*CommitteeContext, error) {
	c := CommitteeContext{
		Context: ctx,
		Messages: req.Messages,
		OutputOpinion: opinion,
		OutputReview: review,		
	}
	c.Leader = d.Members[req.Model]
	if c.Leader == nil {
		return nil, errors.New("leader model not found")
	}
	c.Members = gslice.SliceToKVMapIf(members, func(member string) (string, model.LLM, bool) {
		model := d.Members[member]
		if model == nil {
			return "", nil, false
		}
		return member, model, true
	})
	if len(c.Members) == 0 {
		return nil, errors.New("member model not found")
	}
	return &c, nil
}