package bm25

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

type args struct {
	docsToken  [][]string
	queryToken []string
}

type want struct {
	scores []float64
}

func formatFloat(f float64) string {
	return strconv.FormatFloat(f, 'f', 5, 64)
}
func TestBM25Base(t *testing.T) {
	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "basic",
			args: args{
				docsToken: [][]string{
					{"使命召唤", "这个", "游戏", "好玩吗"},
					{"使命召唤", "武器", "推荐"},
					{"apex", "英雄", "武器", "推荐"},
					{"使命召唤", "攻略"},
					{"使命召唤", "使命召唤", "使命召唤", "使命召唤"},
				},
				queryToken: []string{"使命召唤", "攻略"},
			},
			want: want{
				scores: []float64{0.268312, 0.302228, 0, 2.013078, 0.472418},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bm25 := NewBM25(tt.args.docsToken, 1.2, 0.75)
			scores := bm25.GetAllDocsScore(tt.args.queryToken)
			for idx, exp := range tt.want.scores {
				assert.Equal(t, formatFloat(exp), formatFloat(scores[idx]))
			}
		})
	}

}
