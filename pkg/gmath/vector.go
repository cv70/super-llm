package gmath

import (
	"math"
)

// VectorL2Norm32 模长
func VectorL2Norm32(v []float32) float32 {
	var sum float32
	for _, vi := range v {
		sum += vi * vi
	}
	return float32(math.Sqrt(float64(sum)))
}

// VectorCosine32 余弦距离
func VectorCosine32(v1, v2 []float32) float32 {
	return VectorDot32(v1, v2) / VectorL2Norm32(v1) / VectorL2Norm32(v2)
}

func VectorCosine32Normalize(v1, v2 []float32) float32 {
	return (VectorCosine32(v1, v2) + 1) / 2
}
