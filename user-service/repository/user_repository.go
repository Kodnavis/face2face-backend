package repository

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/Kodnavis/face2face-backend/user-service/model"
	"golang.org/x/crypto/bcrypt"
)

type UserRepo struct {
	Db *sql.DB
}

func (u *UserRepo) Insert(ctx context.Context, user model.User) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("password hashing failed: %w", err)
	}

	user.Password = string(hashedPassword)

	_, err = u.Db.ExecContext(ctx, `
        INSERT INTO users (id, firstname, lastname, login, password) 
        VALUES ($1, $2, $3, $4, $5)`,
		user.ID, user.Firstname, user.Lastname, user.Login, user.Password)

	if err != nil {
		return fmt.Errorf("failed to insert user: %w", err)
	}

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
	rows, err := u.Db.QueryContext(ctx, "SELECT * FROM users ORDER BY id LIMIT $1 OFFSET $2", page.Size, page.Offset)
	if err != nil {
		return FindAllResult{}, fmt.Errorf("failed to query users: %w", err)
	}
	defer rows.Close()

	var users []model.User
	for rows.Next() {
		var data string
		if err := rows.Scan(data); err != nil {
			return FindAllResult{}, fmt.Errorf("failed to scan order: %w", err)
		}

		var user model.User
		if err := json.Unmarshal([]byte(data), &user); err != nil {
			return FindAllResult{}, fmt.Errorf("failed to decode user json: %w", err)
		}

		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		return FindAllResult{}, fmt.Errorf("rows error: %w", err)
	}

	return FindAllResult{
		Users:  users,
		Cursor: page.Offset + uint64(len(users)),
	}, nil
}

var ErrNotExist = errors.New("user does not exist")

func (u *UserRepo) Find(ctx context.Context, id uint64) (model.User, error) {
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

func (u *UserRepo) Update(ctx context.Context, user model.User) error {
	data, err := json.Marshal(user)
	if err != nil {
		return fmt.Errorf("failed to encode order: %w", err)
	}

	_, err = u.Db.ExecContext(ctx, "UPDATE users SET data = $1 WHERE id = $2", data, user.ID)
	if errors.Is(err, sql.ErrNoRows) {
		return ErrNotExist
	} else if err != nil {
		return fmt.Errorf("failed to update user: %w", err)
	}

	return nil
}

func (u *UserRepo) Delete(ctx context.Context, id uint64) error {
	_, err := u.Db.ExecContext(ctx, "DETELE FROM users WHERE id = $1", id)
	if errors.Is(err, sql.ErrNoRows) {
		return ErrNotExist
	} else if err != nil {
		return fmt.Errorf("failed to delete user: %w", err)
	}

	return nil
}
