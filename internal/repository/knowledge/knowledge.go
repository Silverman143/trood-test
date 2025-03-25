package knowrepo

import (
	"context"
	"trood-test/db/postgres"
	"trood-test/internal/domain/models"
)




type KnowledgeRepo struct {
	storage *postgres.Storage
}

func NewRepo(storage *postgres.Storage) *KnowledgeRepo {
	return &KnowledgeRepo{storage: storage}
}

func (k *KnowledgeRepo) StoreKnowledge(ctx context.Context, content string, embedding []float32, metadata map[string]interface{}) (int64, error) {
	return 1, nil
}

func (k *KnowledgeRepo) SearchSimilar(ctx context.Context, queryEmbedding []float32, limit int) ([]models.KnowledgeItem, error) {
	return []models.KnowledgeItem{}, nil
}
