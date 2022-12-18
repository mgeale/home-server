package db

import (
	"context"
	"database/sql"
	"errors"
	"log"
	"time"
)

type User struct {
	ID             int64
	Name           string
	Email          string
	HashedPassword []byte
	Created        time.Time
	Active         bool
}

type UserModel struct {
	DB       *sql.DB
	InfoLog  *log.Logger
	ErrorLog *log.Logger
}

func (m UserModel) GetByEmail(email string) (*User, error) {
	query := `
		SELECT id, name, email, hashed_password, created, active
		FROM users
		WHERE email = ?
		`
	var user User

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := m.DB.QueryRowContext(ctx, query, email).Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.HashedPassword,
		&user.Created,
		&user.Active,
	)

	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrRecordNotFound
		default:
			return nil, err
		}
	}

	return &user, nil
}
