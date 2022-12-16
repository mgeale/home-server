package db

import (
	"database/sql"
	"errors"
	"log"
	"os"
	"time"
)

var (
	ErrInvalidCredentials = errors.New("models: invalid credentials")
	ErrDuplicateEmail     = errors.New("models: duplicate email")
	ErrRecordNotFound     = errors.New("models: record not found")
	ErrEditConflict       = errors.New("models: edit conflict")
)

type Models struct {
	Balances interface {
		Delete(id int) error
		Get(id int) (*Balance, error)
		Insert(name string, balance float64, balanceaud float64, pricebookid int, productid int) (int, error)
		Latest() ([]*Balance, error)
		Update(id int, name string, balance float64, balanceaud float64, pricebookid int, productid int) error
	}
	Transactions interface {
		Delete(id int) error
		Get(id int) (*Transaction, error)
		Insert(name string, amount float64, date string, transactionType string) (int, error)
		Latest() ([]*Transaction, error)
		Update(id int, name string, amount float64, date string, transactionType string) error
	}
	Users interface {
		GetByEmail(email string) (*User, error)
		GetForToken(tokenScope string, tokenPlaintext string) (*User, error)
	}
	Tokens interface {
		DeleteAllForUser(scope string, userID int64) error
		Insert(token *Token) error
		New(userID int64, ttl time.Duration, scope string) (*Token, error)
	}
	Permissions interface {
		AddForUser(userID int64, codes string) error
		GetAllForUser(userID int64) (Permissions, error)
	}
}

func NewModels(db *sql.DB) Models {
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	return Models{
		Balances: &BalanceModel{
			DB:       db,
			InfoLog:  infoLog,
			ErrorLog: errorLog,
		},
		Transactions: &TransactionModel{
			DB:       db,
			InfoLog:  infoLog,
			ErrorLog: errorLog,
		},
		Users: &UserModel{
			DB:       db,
			InfoLog:  infoLog,
			ErrorLog: errorLog,
		},
		Tokens: &TokenModel{
			DB:       db,
			InfoLog:  infoLog,
			ErrorLog: errorLog,
		},
		Permissions: &PermissionModel{
			DB:       db,
			InfoLog:  infoLog,
			ErrorLog: errorLog,
		},
	}
}
