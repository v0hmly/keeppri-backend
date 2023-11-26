package services

import (
	"github.com/v0hmly/keeppri-backend/internal/lib/hash"
	"github.com/v0hmly/keeppri-backend/internal/repository"
	"github.com/v0hmly/keeppri-backend/internal/repository/domain"
)

type AuthServices struct {
	repos *repository.Repository
	hash  hash.PasswordHasher
}

func NewAuthServices(repos *repository.Repository, hash hash.PasswordHasher) *AuthServices {
	return &AuthServices{
		repos: repos,
		hash:  hash,
	}
}

func (s *AuthServices) Register(user *domain.User) (*string, error) {
	userExists, err := s.repos.DB.UserExists(user.Email)
	if err != nil && userExists {
		return nil, err
	}

	hashedPwd, err := s.hash.GenerateHash(user.Password)
	if err != nil {
		return nil, err
	}
	user.Password = hashedPwd

	UserId, err := s.repos.DB.Register(user)
	if err != nil {
		return nil, err
	}
	return UserId, nil
}

func (s *AuthServices) Login(email, password string) (*string, error) {
	s.repos.DB.Login(email, password)

	return nil, nil
}

func (s *AuthServices) Logout(sessionID string) error {
	return nil
}
