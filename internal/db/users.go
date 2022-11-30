package db

import (
	"context"
	"crypto/sha256"
	"database/sql"
	"errors"
	"log"
	"time"

	"github.com/mgeale/homeserver/internal/validator"
	"golang.org/x/crypto/bcrypt"
)

var AnonymousUser = &User{}

type User struct {
	ID             int64
	Name           string
	Email          string
	HashedPassword []byte
	Created        time.Time
	Active         bool
}

func (u *User) IsAnonymous() bool {
	return u == AnonymousUser
}

type UserModel struct {
	DB       *sql.DB
	InfoLog  *log.Logger
	ErrorLog *log.Logger
}

func Matches(hashedPassword []byte, plaintextPassword string) (bool, error) {
	err := bcrypt.CompareHashAndPassword(hashedPassword, []byte(plaintextPassword))
	if err != nil {
		switch {
		case errors.Is(err, bcrypt.ErrMismatchedHashAndPassword):
			return false, nil
		default:
			return false, err
		}
	}

	return true, nil
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

func (m UserModel) GetForToken(tokenScope, tokenPlaintext string) (*User, error) {
	tokenHash := sha256.Sum256([]byte(tokenPlaintext))

	query := `
		SELECT 
			users.id, users.name, users.email, users.hashed_password, users.created
		FROM       users
        INNER JOIN tokens
			ON users.id = tokens.user_id
        WHERE tokens.hash = ?
			AND tokens.scope = ?
			AND tokens.expiry > ?
		`

	args := []interface{}{tokenHash[:], tokenScope, time.Now()}

	var user User

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := m.DB.QueryRowContext(ctx, query, args...).Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.HashedPassword,
		&user.Created,
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

func ValidatePasswordPlaintext(v *validator.Validator, password string) {
	v.Check(password != "", "password", "must be provided")
	v.Check(len(password) >= 8, "password", "must be at least 8 bytes long")
	v.Check(len(password) <= 72, "password", "must not be more than 72 bytes long")
}
