package repositories

import (
	"context"
	"errors"
	"fmt"

	"github.com/Kodnavis/face2face-backend/user-service/internal/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func (u *UserRepository) Insert(user *models.User) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("password hashing failed: %w", err)
	}

	user.Password = string(hashedPassword)

	return u.DB.Create(user).Error
}

type FindAllPage struct {
	Size   uint64
	Offset uint64
}

type FindAllResult struct {
	Users  []models.User
	Cursor uint64
}

func (u *UserRepository) FindAll(ctx context.Context, page FindAllPage) (FindAllResult, error) {
	return FindAllResult{}, nil
}

var ErrNotExist = errors.New("user does not exist")

func (u *UserRepository) Find(ctx context.Context, id uint64) (models.User, error) {
	return models.User{}, nil
}

func (u *UserRepository) Update(ctx context.Context, user models.User) error {
	return nil
}

func (u *UserRepository) Delete(ctx context.Context, id uint64) error {
	return nil
}
