package gmath

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/exp/constraints"
)

func TestGetYByXFunc(t *testing.T) {
	assert.Equal(t, 0.0, GetYByXFunc(0, 0, 1, 1)(0))
	assert.Equal(t, 0.5, GetYByXFunc(0, 0, 1, 1)(0.5))
	assert.Equal(t, 1.0, GetYByXFunc(0, 0, 1, 1)(1))
	assert.Equal(t, 0.0, GetYByXFunc(0, 0, 1, 0)(0))
	assert.Equal(t, 0.0, GetYByXFunc(0, 0, 1, 0)(0.5))
	assert.Equal(t, 0.0, GetYByXFunc(0, 0, 1, 0)(1))
	assert.Equal(t, 1.0, GetYByXFunc(0, 1, 1, 0)(0))
	assert.Equal(t, 0.5, GetYByXFunc(0, 1, 1, 0)(0.5))
	assert.Equal(t, 0.0, GetYByXFunc(0, 1, 1, 0)(1))
}

func TestGreatestCommonDivisor(t *testing.T) {
	type args[T constraints.Integer] struct {
		a T
		b T
	}
	type testCase[T constraints.Integer] struct {
		name string
		args args[T]
		want T
	}
	tests := []testCase[int]{
		{
			name: "",
			args: args[int]{a: 1, b: 1},
			want: 1,
		},
		{
			name: "",
			args: args[int]{a: 319, b: 377},
			want: 29,
		},
		{
			name: "",
			args: args[int]{a: 377, b: 319},
			want: 29,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, GreatestCommonDivisor(tt.args.a, tt.args.b), "GreatestCommonDivisor(%v, %v)", tt.args.a, tt.args.b)
		})
	}
}

func TestLeastCommonMultiple(t *testing.T) {
	type args[T constraints.Integer] struct {
		a T
		b T
	}
	type testCase[T constraints.Integer] struct {
		name string
		args args[T]
		want T
	}
	tests := []testCase[int]{
		{
			name: "",
			args: args[int]{a: 5, b: 6},
			want: 30,
		},
		{
			name: "",
			args: args[int]{a: 15, b: 20},
			want: 60,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, LeastCommonMultiple(tt.args.a, tt.args.b), "LeastCommonMultiple(%v, %v)", tt.args.a, tt.args.b)
		})
	}
}
