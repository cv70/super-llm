package giter

import (
	"iter"
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSum(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		seq  iter.Seq[int]
		want int
	}{
		{
			name: "empty",
			seq:  slices.Values([]int{}),
			want: 0,
		},
		{
			name: "single",
			seq:  slices.Values([]int{1}),
			want: 1,
		},
		{
			name: "multiple",
			seq:  slices.Values([]int{1, 2, 3, 4, 5}),
			want: 15,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, Sum(tt.seq))
		})
	}
}
