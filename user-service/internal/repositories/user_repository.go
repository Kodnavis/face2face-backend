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

type FindAllQueryParams struct {
	Size   int `form:"size"`
	Offset int `form:"offset"`
}

func (u *UserRepository) FindAll(params FindAllQueryParams) ([]models.User, error) {
	var users []models.User

	result := u.DB.Limit(params.Size).Offset(params.Offset).Find(&users)
	if result.Error != nil {
		return []models.User{}, fmt.Errorf("listing users failed: %w", result.Error)
	}

	return users, nil
}

var ErrNotExist = errors.New("user does not exist")

func (u *UserRepository) Find(ctx context.Context, id uint64) (models.User, error) {
	// TODO
	return models.User{}, nil
}

func (u *UserRepository) Update(ctx context.Context, user models.User) error {
	// TODO
	return nil
}

func (u *UserRepository) Delete(ctx context.Context, id uint64) error {
	// TODO
	return nil
}
