//go:build amd64

package gmath

func vectorDot32(x, y []float32) (sum float32)

func VectorDot32(x, y []float32) (sum float32) {
	if len(x) != len(y) {
		panic("两向量长度不同")
	}

	return vectorDot32(x, y)
}
