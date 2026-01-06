//go:build !amd64

package gmath

// VectorDot32 点积
func VectorDot32(v1, v2 []float32) float32 {
	if len(v1) != len(v2) {
		panic("两向量长度不同")
	}

	var dot float32
	for i, v1I := range v1 {
		v2I := v2[i]
		dot += v1I * v2I
	}
	return dot
}
