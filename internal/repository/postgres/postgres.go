package postgres

import (
	"fmt"

	"github.com/v0hmly/keeppri-backend/internal/config"
	"github.com/v0hmly/keeppri-backend/internal/repository/domain"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type UserRepository interface {
	UserExists(email string) (bool, error)
	Login(email, password string) (*string, error)
	Register(user *domain.User) (*string, error)
}

type DBConn struct {
	db *gorm.DB
	UserRepository
}

func NewDB(cfg *config.PostgresConfig) (*DBConn, error) {
	op := "repository.postgres.NewDB"

	conn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.Db)

	db, err := gorm.Open(postgres.Open(conn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("%s: db connection failed: %w", op, err)
	}

	return &DBConn{db: db}, nil
}
