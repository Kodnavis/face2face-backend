package repository

import (
	"context"
	"database/sql"
	"encoding/json"
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
