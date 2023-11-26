package redis

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
	"github.com/v0hmly/keeppri-backend/internal/config"
	"github.com/v0hmly/keeppri-backend/internal/repository/domain"
)

type SessionsRepository interface {
	SetSession(session *domain.Session) error
	GetSession(sessionID string) (*domain.Session, error)
	DelSession(sessionID string) error
}

type Redis struct {
	client *redis.Client
}

func NewRedis(cfg config.RedisConfig) (*Redis, error) {
	const op = "repository.redis.NewRedis"
	client := redis.NewClient(&redis.Options{
		Addr: cfg.Host + ":" + fmt.Sprint(cfg.Port),
	})

	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &Redis{client: client}, nil
}
