package tfidf

import (
	"math"
)

type TFIDF struct {
	N   int
	df  map[string]int
	idf map[string]float64
}

func NewTFIDF(docsTokens [][]string) *TFIDF {
	tfidf := &TFIDF{
		N:   len(docsTokens),
		df:  make(map[string]int),
		idf: make(map[string]float64),
	}

	for _, docTokens := range docsTokens {
		seen := make(map[string]bool)
		for _, token := range docTokens {
			if !seen[token] {
				tfidf.df[token]++
				seen[token] = true
			}
		}
	}

	for token, df := range tfidf.df {
		tfidf.idf[token] = math.Log(1 + (float64(tfidf.N)-float64(df)+0.5)/(float64(df)+0.5))
	}

	return tfidf
}

func (tfidf *TFIDF) GetIDF(token string) float64 {
	return tfidf.idf[token]
}
