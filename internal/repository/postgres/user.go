package postgres

import (
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/v0hmly/keeppri-backend/internal/repository/domain"
	"gorm.io/gorm"
)

func (d *DBConn) Register(user *domain.User) (*string, error) {
	op := "repository.postgres.Register"

	userID := uuid.New()
	user.ID = userID

	req := d.db.Create(user)
	if req.Error != nil {
		if errors.Is(req.Error, gorm.ErrDuplicatedKey) {
			return nil, ErrUserExists
		}
		return nil, fmt.Errorf("%s: %v", op, req.Error)
	}

	userIDStr := userID.String()
	return &userIDStr, nil
}

func (d *DBConn) GetUserDataByEmail(email string) (*domain.User, error) {
	op := "repository.postgres.GetUserDataByEmail"

	user := &domain.User{}
	req := d.db.Where("email = ?", email).First(&user)
	if req.RowsAffected == 0 {
		return nil, ErrUserNotFound
	}
	if req.Error != nil {
		return nil, fmt.Errorf("%s: %w", op, req.Error)
	}
	return user, nil
}

var (
	ErrUserExists   = errors.New("user already exists")
	ErrUserNotFound = errors.New("user not found")
)
