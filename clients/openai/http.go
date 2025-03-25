package openai

import (
	"context"
	"net/http"
	"time"
	"trood-test/env"
)

type client struct {
	client     http.Client
	baseURL    string
	apiKey     string
	timeout    time.Duration
	maxRetries int
}

func NewClient(cfg *env.OpenaiClient) *client {
	return &client{
		client:     http.Client{},
		baseURL:    "https://api.openai.com/v1",
		apiKey:     cfg.APIKey,
		timeout:    cfg.Timeout,
		maxRetries: cfg.RetriesCount,
	}
}

func (c *client) GenerateEmbedding(ctx context.Context, text string) ([]float32, error) {
	// Sends a query to the desired model to get a vector representation of the text 
	return []float32{}, nil
}
func (c *client) GenerateAnswer(ctx context.Context, prompt string, info []string) (string, error){

	// use info as context 
	return "", nil
}
