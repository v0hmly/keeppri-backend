package services

import (
	"log/slog"

	"github.com/v0hmly/keeppri-backend/internal/lib/hash"
	"github.com/v0hmly/keeppri-backend/internal/lib/token"
	"github.com/v0hmly/keeppri-backend/internal/repository"
	"github.com/v0hmly/keeppri-backend/internal/repository/domain"
)

type (
	AuthService interface {
		Register(user *domain.User) (*string, error)
		Login(email, password string) (*string, error)
		Logout(sessionToken string) error
	}

	Services struct {
		AuthService AuthService
	}

	Deps struct {
		Logger       *slog.Logger
		Repos        *repository.Repository
		Hash         hash.PasswordHasher
		TokenManager token.TokenManager
	}
)

func NewServices(deps Deps) *Services {
	return &Services{
		AuthService: NewAuthServices(deps.Logger, deps.Repos, deps.Hash, deps.TokenManager),
	}
}
