package sdk

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

type EmbeddingClient interface {
	Embedding(ss ...string) ([][]float32, error)
}

type AnythingEmbeddingClient struct {
	Type string
	URL string
}

var _ EmbeddingClient = &AnythingEmbeddingClient{}

type EmbeddingResult struct {
	Embeddings [][]float32 `json:"embddings"`
}

func (c *AnythingEmbeddingClient) Embedding(ss ...string) ([][]float32, error) {
	if len(ss) == 0 {
		return [][]float32{}, nil
	}
	req := map[string]any{
		"ss": ss,
	}
	reqData, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	reqReader := bytes.NewBuffer(reqData)

	resp, err := http.Post(c.URL, "application/json", reqReader)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	respData, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	
	embeddingResult := EmbeddingResult{}
	err = json.Unmarshal(respData, &embeddingResult)
	if err != nil {
		return nil, err
	}

	return embeddingResult.Embeddings, err
}
