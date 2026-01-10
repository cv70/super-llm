package committee

import (
	"context"

	"github.com/cv70/pkgo/llm"

	"github.com/cv70/pkgo/gslice"

	"github.com/pkg/errors"
)

type CommitteeContext struct {
	context.Context
	Request  *llm.ChatCompletionRequest
	Messages []*llm.ChatMessage
	Leader   *llm.OpenAIModel
	Members  map[string]*llm.OpenAIModel

	Opinions       map[string]string
	Reviews        map[string][]string
	MessageSummary string

	OutputOpinion bool
	OutputReview  bool
}

func (d *CommitteeDomain) BuildCommitteeContext(ctx context.Context, req *llm.ChatCompletionRequest, members []string, opinion, review bool) (*CommitteeContext, error) {
	c := CommitteeContext{
		Context:       ctx,
		Request:       req,
		Messages:      req.Messages,
		OutputOpinion: opinion,
		OutputReview:  review,
	}
	c.Leader = d.Members[req.Model]
	if c.Leader == nil {
		return nil, errors.New("leader model not found")
	}
	if len(members) == 0 {
		c.Members = d.Members
	} else {
		c.Members = gslice.SliceToMapIf(members, func(member string) (string, *llm.OpenAIModel, bool) {
			model := d.Members[member]
			if model == nil {
				return "", nil, false
			}
			return member, model, true
		})
	}
	if len(c.Members) == 0 {
		return nil, errors.New("member model not found")
	}
	return &c, nil
}
