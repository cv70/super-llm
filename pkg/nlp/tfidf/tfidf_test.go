package tfidf

import (
	"math"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func formatFloat(f float64) string {
	return strconv.FormatFloat(f, 'f', 5, 64)
}

func TestNewTFIDF(t *testing.T) {
	docs := [][]string{
		{"hello", "world"},
		{"hello", "golang"},
		{"bm25", "test", "example"},
	}
	tfidf := NewTFIDF(docs)
	if tfidf.N != 3 {
		t.Errorf("Expected N = 3, got %d", tfidf.N)
	}
	expectedDF := map[string]int{
		"hello":   2,
		"world":   1,
		"golang":  1,
		"bm25":    1,
		"test":    1,
		"example": 1,
	}

	for token, expected := range expectedDF {
		assert.Equal(t, expected, tfidf.df[token])
	}

	expectedIDF := map[string]float64{
		"hello":   math.Log(1 + (3-2+0.5)/(2+0.5)), // log(1 + (3-2+0.5) / (2+0.5))
		"world":   math.Log(1 + (3-1+0.5)/(1+0.5)),
		"golang":  math.Log(1 + (3-1+0.5)/(1+0.5)),
		"bm25":    math.Log(1 + (3-1+0.5)/(1+0.5)),
		"test":    math.Log(1 + (3-1+0.5)/(1+0.5)),
		"example": math.Log(1 + (3-1+0.5)/(1+0.5)),
	}

	for token, expected := range expectedIDF {
		assert.Equal(t, formatFloat(expected), formatFloat(tfidf.idf[token]))
	}
}

func TestGetIDF(t *testing.T) {
	docs := [][]string{
		{"hello", "world"},
		{"hello", "golang"},
	}

	var (
		tfidf = NewTFIDF(docs)
		n     = 2.
		df    = 2.
	)
	expectedIDFHello := math.Log(1 + (n-df+0.5)/(df+0.5))
	idfHello := tfidf.GetIDF("hello")

	assert.Equal(t, formatFloat(expectedIDFHello), formatFloat(idfHello))

	idfUnknown := tfidf.GetIDF("unknown")
	if idfUnknown != 0.0 {
		t.Errorf("GetIDF(unknown): expected 0.0, got %.4f", idfUnknown)
	}
}
