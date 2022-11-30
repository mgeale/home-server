package mock

import (
	"time"

	"github.com/mgeale/homeserver/internal/db"
)

var mockUser = &db.User{
	ID:      1,
	Name:    "Alice",
	Email:   "alice@example.com",
	Created: time.Now(),
	Active:  true,
}

type UserModel struct{}

func (m *UserModel) Insert(name, email, password string) error {
	switch email {
	case "dupe@example.com":
		return db.ErrDuplicateEmail
	default:
		return nil
	}
}

func (m *UserModel) Authenticate(email, password string) error {
	switch email {
	case "alice@example.com":
		return nil
	default:
		return db.ErrInvalidCredentials
	}
}

func (m *UserModel) Get(id int) (*db.User, error) {
	switch id {
	case 1:
		return mockUser, nil
	default:
		return nil, db.ErrRecordNotFound
	}
}
