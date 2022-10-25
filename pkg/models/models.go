package models

import (
	"errors"
	"time"
)

var (
	ErrNoRecord           = errors.New("models: no matching record found")
	ErrInvalidCredentials = errors.New("models: invalid credentials")
	ErrDuplicateEmail     = errors.New("models: duplicate email")
)

type Balance struct {
	ID          int
	Name        string
	Balance     float32
	BalanceAUD  float32
	PricebookID int
	ProductID   int
	Created     time.Time
}

type Transaction struct {
	ID      int
	Name    string
	Amount  float32
	Date    string
	Type    string
	Created time.Time
}

type User struct {
	ID             int
	Name           string
	Email          string
	HashedPassword []byte
	Created        time.Time
	Active         bool
}
