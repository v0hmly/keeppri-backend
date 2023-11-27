package domain

import (
	"time"

	"github.com/google/uuid"
)

type Session struct {
	SessionToken string        `json:"session_token"`
	UserID       uuid.UUID     `json:"user_id"`
	ExpireAt     time.Duration `json:"expire_at"`
}
