package db

import (
	"database/sql"
	"errors"
	"log"
	"os"
)

var (
	ErrInvalidCredentials = errors.New("models: invalid credentials")
	ErrDuplicateEmail     = errors.New("models: duplicate email")
	ErrRecordNotFound     = errors.New("models: record not found")
	ErrEditConflict       = errors.New("models: edit conflict")
)

type Models struct {
	Balances     BalanceModel
	Transactions TransactionModel
	Users        UserModel
	Tokens       TokenModel
	Permissions  PermissionModel
}

func NewModels(db *sql.DB) Models {
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	return Models{
		Balances: BalanceModel{
			DB:       db,
			InfoLog:  infoLog,
			ErrorLog: errorLog,
		},
		Transactions: TransactionModel{
			DB:       db,
			InfoLog:  infoLog,
			ErrorLog: errorLog,
		},
		Users: UserModel{
			DB:       db,
			InfoLog:  infoLog,
			ErrorLog: errorLog,
		},
		Tokens: TokenModel{
			DB:       db,
			InfoLog:  infoLog,
			ErrorLog: errorLog,
		},
		Permissions: PermissionModel{
			DB:       db,
			InfoLog:  infoLog,
			ErrorLog: errorLog,
		},
	}
}
