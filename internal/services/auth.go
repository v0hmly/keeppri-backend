package services

import (
	"errors"
	"fmt"
	"log/slog"

	"github.com/v0hmly/keeppri-backend/internal/lib/hash"
	log "github.com/v0hmly/keeppri-backend/internal/lib/logger"
	"github.com/v0hmly/keeppri-backend/internal/repository"
	"github.com/v0hmly/keeppri-backend/internal/repository/domain"
)

type AuthServices struct {
	logger *slog.Logger
	repos  *repository.Repository
	hash   hash.PasswordHasher
}

func NewAuthServices(logger *slog.Logger, repos *repository.Repository, hash hash.PasswordHasher) *AuthServices {
	return &AuthServices{
		logger: logger,
		repos:  repos,
		hash:   hash,
	}
}

func (s *AuthServices) Register(user *domain.User) (*string, error) {
	op := "services.AuthServices.Register"

	logger := s.logger.With(
		slog.String("op", op),
		slog.String("email", user.Email),
	)

	logger.Debug("registering user")

	userExists, err := s.repos.DB.UserExists(user.Email)
	if err != nil {
		logger.Error("failed to detect that the user exists", log.Err(err))
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	if userExists {
		logger.Error("user already exists")
		return nil, ErrUserExists
	}

	hashedPwd, err := s.hash.GenerateHash(user.Password)
	if err != nil {
		logger.Error("failed to generate password hash", log.Err(err))
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	user.Password = hashedPwd

	UserId, err := s.repos.DB.Register(user)
	if err != nil {
		logger.Error("failed to register user", log.Err(err))
		return nil, fmt.Errorf("%s: %w", op, err)
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

var (
	ErrUserExists   = errors.New("user already exists")
	ErrUserNotFound = errors.New("user not found")
)
