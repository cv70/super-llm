package gmath

import (
	"testing"
)

var (
	v1, v2 []float32
)

func init() {
	v1 = make([]float32, 512)
	v2 = make([]float32, 512)
}

func BenchmarkVectorDot32(b *testing.B) {
	for i := 0; i < b.N; i++ {
		VectorDot32(v1, v2)
	}
}
