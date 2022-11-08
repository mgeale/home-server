package app

import (
	"context"
	"errors"
	"log"
	"sync"

	"github.com/mgeale/homeserver/internal/db"
)

type config struct {
	Port int
	Env  string
	db   struct {
		dsn          string
		maxOpenConns int
		maxIdleConns int
		maxIdleTime  string
	}
}

type Application struct {
	Config   config
	ErrorLog *log.Logger
	InfoLog  *log.Logger
	Models   *db.Models
	Wg       sync.WaitGroup
}

func (app *Application) CreateBalance(ctx context.Context, balance *db.Balance) (int, error) {
	int, err := app.Models.Balances.Insert(balance.Name, balance.Balance, balance.BalanceAUD, balance.PricebookID, balance.ProductID)
	if err != nil {
		return 0, err
	}

	return int, nil
}

func (app *Application) CreateTransaction(ctx context.Context, transaction *db.Transaction) (int, error) {
	int, err := app.Models.Transactions.Insert(transaction.Name, transaction.Amount, transaction.Date, transaction.Type)
	if err != nil {
		return 0, err
	}

	return int, nil
}

func (app *Application) UpdateBalance(ctx context.Context, balance *db.Balance) (int, error) {
	// if b == nil || b.Name == "" || b.Balance == 0 || b.BalanceAUD == 0 || b.PricebookID == 0 || b.ProductID == 0 {
	// 	app.clientError(w, http.StatusBadRequest)
	// 	return
	// }

	err := app.Models.Balances.Update(balance.ID, balance.Name, balance.Balance, balance.BalanceAUD, balance.PricebookID, balance.ProductID)
	if err != nil {
		if errors.Is(err, db.ErrNoRecord) {
			//TODO: not found
			return 0, err
		} else {
			//TODO: server errr
			return 0, err
		}
	}

	return 1, nil
}

func (app *Application) UpdateTransaction(ctx context.Context, transaction *db.Transaction) (int, error) {
	// 	if t == nil || t.Name == "" || t.Amount == 0 || t.Date == "" || t.Type == "" {
	// 		app.clientError(w, http.StatusBadRequest)
	// 		return
	// 	}

	int, err := app.Models.Transactions.Insert(transaction.Name, transaction.Amount, transaction.Date, transaction.Type)
	if err != nil {
		if errors.Is(err, db.ErrNoRecord) {
			//TODO: not found
			return 0, err
		} else {
			//TODO: server errr
			return 0, err
		}
	}

	return int, nil
}

func (app *Application) DeleteBalance(ctx context.Context, id int) (int, error) {
	err := app.Models.Balances.Delete(id)
	if err != nil {
		if errors.Is(err, db.ErrNoRecord) {
			//TODO: not found
			return 0, err
		} else {
			//TODO: server errr
			return 0, err
		}
	}
	return 1, nil
}

func (app *Application) DeleteTransaction(ctx context.Context, id int) (int, error) {
	err := app.Models.Transactions.Delete(id)
	if err != nil {
		if errors.Is(err, db.ErrNoRecord) {
			//TODO: not found
			return 0, err
		} else {
			//TODO: server errr
			return 0, err
		}
	}
	return 1, nil
}

func (app *Application) GetLatestBalances() ([]*db.Balance, error) {
	b, err := app.Models.Balances.Latest()
	if err != nil {
		if errors.Is(err, db.ErrNoRecord) {
			//TODO: not found
			return nil, err
		} else {
			//TODO: server errr
			return nil, err
		}
	}

	return b, nil
}

func (app *Application) GetLatestTransactions() ([]*db.Transaction, error) {
	b, err := app.Models.Transactions.Latest()
	if err != nil {
		if errors.Is(err, db.ErrNoRecord) {
			//TODO: not found
			return nil, err
		} else {
			//TODO: server errr
			return nil, err
		}
	}

	return b, nil
}
