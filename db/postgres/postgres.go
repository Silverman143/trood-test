package postgres

import (
	"fmt"
	"trood-test/env"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // Явно импортируем драйвер PostgreSQL
)

type Storage struct {
	DB *sqlx.DB
}

func New(cfg *env.PgSql) (*Storage, error) {
	const op = "repository.postgres.postgres.NewDBConnection"

	var connStr string
	if cfg.URI == ""{
		connStr = fmt.Sprintf("user=%s dbname=%s sslmode=%s password=%s host=%s",
        cfg.User, cfg.DbName, cfg.SSLMode, cfg.Password, cfg.Host)
	} else {
		connStr = cfg.URI
	}

	db, err := sqlx.Connect("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	return &Storage{DB: db}, nil
}

func (s *Storage) Stop() error {
	return s.DB.Close()
}