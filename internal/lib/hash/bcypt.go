package hash

import (
	"errors"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type BcryptPasswordHasher struct {
}

func NewBcryptPasswordHasher() *BcryptPasswordHasher {
	return &BcryptPasswordHasher{}
}

func (s *BcryptPasswordHasher) GenerateHash(password string) (string, error) {
	op := "hash.bcrypt.GenerateHash"

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("%s: %w", op, err)
	}
	return string(hash), nil
}

func (s *BcryptPasswordHasher) CompareHashAndPassword(hash, password string) (bool, error) {
	op := "hash.bcrypt.CompareHashAndPassword"

	if err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)); err != nil {
		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			return false, nil
		}

		return false, fmt.Errorf("%s: %w", op, err)
	}
	return true, nil
}
