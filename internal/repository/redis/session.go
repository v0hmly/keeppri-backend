package redis

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/go-redis/redis/v8"
	"github.com/v0hmly/keeppri-backend/internal/repository/domain"
)

func (r *Redis) SetSession(session *domain.Session) error {
	op := "repository.redis.session.SetSession"

	data, err := json.Marshal(session)
	if err != nil {
		return fmt.Errorf("%s: unmarshal error: %w", op, err)
	}

	if err := r.client.Set(context.Background(), session.SessionID, data, session.ExpireAt).Err(); err != nil {
		return fmt.Errorf("%s: set error: %w", op, err)
	}

	return nil
}

func (r *Redis) GetSession(sessionID string) (*domain.Session, error) {
	op := "repository.redis.session.GetSession"

	data, err := r.client.Get(context.Background(), sessionID).Bytes()
	if err == redis.Nil {
		return nil, fmt.Errorf("%s: key not found: %w", op, err)
	} else if err != nil {
		return nil, fmt.Errorf("%s: get error: %w", op, err)
	}

	var session *domain.Session

	if err := json.Unmarshal(data, session); err != nil {
		return nil, fmt.Errorf("%s: unmarshal error: %w", op, err)
	}

	if err := r.client.Expire(context.Background(), sessionID, session.ExpireAt).Err(); err != nil {
		return nil, fmt.Errorf("%s: expire error: %w", op, err)
	}

	return session, nil
}

func (r *Redis) DelSession(sessionID string) error {
	op := "repository.redis.session.DelSession"

	if err := r.client.Del(context.Background(), sessionID).Err(); err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	return nil

}
