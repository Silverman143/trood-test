package models

import (
	"time"

	"github.com/lib/pq"
)

type KnowledgeItem struct {
	ID        int64             `db:"id" json:"id"`
	Content   string            `db:"content" json:"content"`
	Metadata  map[string]interface{} `db:"metadata" json:"metadata"`
	Embedding pq.Float32Array   `db:"embedding" json:"-"`
	CreatedAt time.Time         `db:"created_at" json:"created_at"`
	UpdatedAt time.Time         `db:"updated_at" json:"updated_at"`
}

func ExtractContents(items []KnowledgeItem) []string {
	contents := make([]string, len(items))
	
	for i, item := range items {
		contents[i] = item.Content
	}
	
	return contents
}