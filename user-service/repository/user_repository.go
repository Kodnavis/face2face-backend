package repository

import (
	"context"
	"errors"

	"github.com/Kodnavis/face2face-backend/user-service/model"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserRepo struct {
	Db *gorm.DB
}

func (u *UserRepo) Insert(ctx *gin.Context, user model.User) error {
	return nil
}

type FindAllPage struct {
	Size   uint64
	Offset uint64
}

type FindAllResult struct {
	Users  []model.User
	Cursor uint64
}

func (u *UserRepo) FindAll(ctx context.Context, page FindAllPage) (FindAllResult, error) {
	return FindAllResult{}, nil
}

var ErrNotExist = errors.New("user does not exist")

func (u *UserRepo) Find(ctx context.Context, id uint64) (model.User, error) {
	return model.User{}, nil
}

func (u *UserRepo) Update(ctx context.Context, user model.User) error {
	return nil
}

func (u *UserRepo) Delete(ctx context.Context, id uint64) error {
	return nil
}
