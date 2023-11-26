package domain

import (
	"time"

	"github.com/google/uuid"
)

type Session struct {
	SessionID string        `json:"session_id"`
	UserID    uuid.UUID     `json:"user_id"`
	ExpireAt  time.Duration `json:"expire_at"`
}
