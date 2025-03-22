package repository

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/Kodnavis/face2face-backend/user-service/model"
)

type UserRepo struct {
	Db *sql.DB
}

func (u *UserRepo) Insert(ctx context.Context, user model.User) error {
	data, err := json.Marshal(user)
	if err != nil {
		return fmt.Errorf("failed to encode order: %w", err)
	}

	_, err = u.Db.ExecContext(ctx, "INSERT INTO users (data) VALUE ($1)", data)
	if err != nil {
		return fmt.Errorf("failed to insert user: %w", err)
	}

	return nil
}

var ErrNotExist = errors.New("user does not exist")

func (u *UserRepo) FindByID(ctx context.Context, id uint64) (model.User, error) {
	var data string
	err := u.Db.QueryRowContext(ctx, "SELECT data from users WHERE $1", id).Scan(&data)
	if errors.Is(err, sql.ErrNoRows) {
		return model.User{}, ErrNotExist
	} else if err != nil {
		return model.User{}, fmt.Errorf("failed to decode order json: %w", err)
	}

	var user model.User
	err = json.Unmarshal([]byte(data), &user)
	if err != nil {
		return model.User{}, fmt.Errorf("failed to decode user json: %w", err)
	}

	return user, nil
}
