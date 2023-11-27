package services

import (
	"errors"
	"fmt"
	"log/slog"
	"time"

	"github.com/google/uuid"
	"github.com/v0hmly/keeppri-backend/internal/lib/hash"
	log "github.com/v0hmly/keeppri-backend/internal/lib/logger"
	"github.com/v0hmly/keeppri-backend/internal/repository"
	"github.com/v0hmly/keeppri-backend/internal/repository/domain"
	"github.com/v0hmly/keeppri-backend/internal/repository/postgres"
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
	op := "services.AuthServices.Login"

	logger := s.logger.With(
		slog.String("op", op),
		slog.String("email", email),
	)

	logger.Debug("registering user")

	user, err := s.repos.DB.GetUserDataByEmail(email)
	if err != nil {
		if errors.Is(err, postgres.ErrUserNotFound) {
			return nil, ErrLoginCredsInvalid
		}
		logger.Error("failed to login user", log.Err(err))
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	arePasswordEquals, err := s.hash.CompareHashAndPassword(user.Password, password)
	if err != nil {
		logger.Error("failed to login user", log.Err(err))
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	if !arePasswordEquals {
		return nil, ErrLoginCredsInvalid
	}

	session := domain.Session{
		SessionToken: uuid.New().String(),
		ExpireAt:     30 * 24 * time.Hour,
		UserID:       user.ID,
	}
	if err := s.repos.Redis.SetSession(&session); err != nil {
		logger.Error("failed to login user", log.Err(err))
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &session.SessionToken, nil
}

func (s *AuthServices) Logout(sessionToken string) error {
	op := "services.AuthServices.Logout"

	logger := s.logger.With(
		slog.String("op", op),
		slog.String("session_token", sessionToken),
	)

	logger.Debug("logout user")

	err := s.repos.Redis.DelSession(sessionToken)
	if err != nil {
		logger.Error("failed to logout user", log.Err(err))
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

var (
	ErrLoginCredsInvalid = errors.New("invalid login credentials")
)
