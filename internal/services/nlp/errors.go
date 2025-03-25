package nlp

import (
	"errors"
)

var (
	ErrEmbeddingGeneration = errors.New("failed to generate embedding")
	ErrKnowledgeSearch = errors.New("failed to search similar knowledge items")
	ErrAnswerGeneration = errors.New("failed to generate answer")
	ErrSendEvent = errors.New("failed to send event")
	ErrUnresolvedIntent = errors.New("unresolved intent")
)