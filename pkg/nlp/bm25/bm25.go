// Package bm25 参考公式 https://www.elastic.co/cn/blog/practical-bm25-part-2-the-bm25-algorithm-and-its-variables
package bm25

import (
	"super-llm/pkg/nlp/tfidf"
)

type BM25 struct {
	k1             float64
	b              float64
	avgDocTokenLen float64
	docs           [][]string
	tfidf          *tfidf.TFIDF
}

func NewBM25(docsTokens [][]string, k1, b float64) *BM25 {
	if len(docsTokens) == 0 {
		return nil
	}
	tfidfModel := tfidf.NewTFIDF(docsTokens)

	var totalLen int
	for _, docTokens := range docsTokens {
		totalLen += len(docTokens)
	}
	return &BM25{
		k1:             k1,
		b:              b,
		avgDocTokenLen: float64(totalLen) / float64(len(docsTokens)),
		docs:           docsTokens,
		tfidf:          tfidfModel,
	}
}

// Score 计算单个文档的 BM25 评分
func (bm25 *BM25) Score(queryTokens []string, docTokens []string) float64 {
	if len(queryTokens) == 0 || len(docTokens) == 0 {
		return 0
	}
	var (
		docLen = len(docTokens)
		tf     = make(map[string]float64)
	)

	for _, token := range docTokens {
		tf[token]++
	}

	var score float64
	for _, token := range queryTokens {
		idf := bm25.tfidf.GetIDF(token)
		tfVal := tf[token]
		numerator := tfVal * (bm25.k1 + 1)
		denominator := tfVal + bm25.k1*(1-bm25.b+bm25.b*float64(docLen)/bm25.avgDocTokenLen)
		score += idf * (numerator / denominator)
	}
	return score
}

func (bm25 *BM25) GetAllDocsScore(queryTokens []string) []float64 {
	scores := make([]float64, 0, len(queryTokens))
	for _, docTokens := range bm25.docs {
		scores = append(scores, bm25.Score(queryTokens, docTokens))
	}
	return scores
}
