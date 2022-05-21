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
	ID         int
	Title      string
	Balance    int
	BalanceAud int
	PriceBook  int
	Product    int
	Created    time.Time
}
