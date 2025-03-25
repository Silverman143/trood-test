package repository

import (
	"context"
	"trood-test/db/postgres"
	"trood-test/internal/domain/models"
	knowrepo "trood-test/internal/repository/knowledge"
)




type IKnowledgeRepo interface{
	StoreKnowledge(ctx context.Context, content string, embedding []float32, metadata map[string]interface{}) (int64, error)
	SearchSimilar(ctx context.Context, queryEmbedding []float32, limit int) ([]models.KnowledgeItem, error)
}

type Repository struct {
	IKnowledgeRepo
}

func New(storage *postgres.Storage) *Repository {
    return &Repository{
		IKnowledgeRepo: knowrepo.NewRepo(storage),
    }
}