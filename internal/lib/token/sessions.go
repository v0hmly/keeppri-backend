package token

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
)

type SessionTokenGenerator struct {
	tokenSize int
}

func NewSessionTokenGenerator(tokenSize int) *SessionTokenGenerator {
	return &SessionTokenGenerator{tokenSize: tokenSize}
}

func (g *SessionTokenGenerator) GenerateToken() (string, error) {
	op := "token.NewSessionTokenGenerator.GenerateToken"

	token := make([]byte, g.tokenSize)

	_, err := rand.Read(token)
	if err != nil {
		return "", fmt.Errorf("%s: %w", op, err)
	}

	return hex.EncodeToString(token), nil
}
