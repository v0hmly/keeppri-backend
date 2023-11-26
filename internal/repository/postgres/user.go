package postgres

import (
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/v0hmly/keeppri-backend/internal/repository/domain"
)

func (d *DBConn) UserExists(email string) (bool, error) {
	op := "repository.postgres.UserExists"
	user := &domain.User{}
	req := d.db.Where("email = ?", email).First(&user)
	if req.RowsAffected == 0 {
		return false, fmt.Errorf("%s: %w", op, ErrUserNotFound)
	}
	return true, nil
}

func (d *DBConn) Register(user *domain.User) (*string, error) {
	op := "repository.postgres.Register"

	userID := uuid.New()
	user.ID = userID

	req := d.db.Create(user)
	if req.RowsAffected == 0 {
		return nil, fmt.Errorf("%s: %v", op, req.Error)
	}

	userIDStr := userID.String()
	return &userIDStr, nil
}

func (d *DBConn) Login(email, password string) (*domain.User, error) {

	user := &domain.User{}
	req := d.db.Select("users").Where("email = ?", email).
		InnerJoins("passwords").First(&user)
	if req.RowsAffected != 0 {
		return nil, errors.New("user+password not found")
	}

	return user, nil
}

var (
	ErrUserExists   = errors.New("user already exists")
	ErrUserNotFound = errors.New("user not found")
	ErrAppNotFound  = errors.New("app not found")
)
