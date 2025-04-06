package repositories

import (
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

func (u *UserRepository) FindOne(login string) (models.User, error) {
	var user models.User

	result := u.DB.Where("login = ?", login).First(&user)
	if result.Error != nil {
		return user, ErrNotExist
	}

	return user, nil
}

func (u *UserRepository) Update(login string, user *models.User) error {
	return u.DB.Where("login = ?", login).Updates(user).Error
}

func (u *UserRepository) Delete(login string) error {
	result := u.DB.Where("login = ?", login).Delete(&models.User{})
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return ErrNotExist
	}

	return nil
}
