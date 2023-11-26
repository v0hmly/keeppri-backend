package repository

import (
	"fmt"

	"github.com/v0hmly/keeppri-backend/internal/config"
	"github.com/v0hmly/keeppri-backend/internal/repository/postgres"
	"github.com/v0hmly/keeppri-backend/internal/repository/redis"
)

type Repository struct {
	DB    *postgres.DBConn
	Redis *redis.Redis
}

func New(cfg *config.Config) (*Repository, error) {
	const op = "repository.repository.New"

	db, err := postgres.NewDB(&cfg.Postgres)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	rd, err := redis.NewRedis(cfg.Redis)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	return &Repository{DB: db, Redis: rd}, nil
}
