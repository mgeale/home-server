package app

import (
	"context"
	"sync"

	"github.com/mgeale/homeserver/internal/db"
	"github.com/mgeale/homeserver/internal/jsonlog"
)

type Config struct {
	Port int
	Env  string
	Db   struct {
		Dsn          string
		MaxOpenConns int
		MaxIdleConns int
		MaxIdleTime  string
	}
	Limiter struct {
		Rps     float64
		Burst   int
		Enabled bool
	}
}

type Application struct {
	Config Config
	Logger *jsonlog.Logger
	Models db.Models
	Wg     sync.WaitGroup
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
	err := app.Models.Balances.Update(balance.ID, balance.Name, balance.Balance, balance.BalanceAUD, balance.PricebookID, balance.ProductID)
	if err != nil {
		return 0, err
	}

	return balance.ID, nil
}

func (app *Application) UpdateTransaction(ctx context.Context, transaction *db.Transaction) (int, error) {
	err := app.Models.Transactions.Update(transaction.ID, transaction.Name, transaction.Amount, transaction.Date, transaction.Type)
	if err != nil {
		return 0, err
	}

	return transaction.ID, nil
}

func (app *Application) DeleteBalance(ctx context.Context, id int) (int, error) {
	err := app.Models.Balances.Delete(id)
	if err != nil {
		return 0, err
	}
	return 1, nil
}

func (app *Application) DeleteTransaction(ctx context.Context, id int) (int, error) {
	err := app.Models.Transactions.Delete(id)
	if err != nil {
		return 0, err
	}
	return 1, nil
}

func (app *Application) GetLatestBalances() ([]*db.Balance, error) {
	b, err := app.Models.Balances.Latest()
	if err != nil {
		return nil, err
	}

	return b, nil
}

func (app *Application) GetBalances(ctx context.Context, query *db.Query) ([]*db.Balance, error) {
	b, err := app.Models.Balances.Get(query)
	if err != nil {
		return nil, err
	}

	return b, nil
}

func (app *Application) GetLatestTransactions() ([]*db.Transaction, error) {
	b, err := app.Models.Transactions.Latest()
	if err != nil {
		return nil, err
	}

	return b, nil
}

func (app *Application) GetTransactions(ctx context.Context, query *db.Query) ([]*db.Transaction, error) {
	b, err := app.Models.Transactions.Get(query)
	if err != nil {
		return nil, err
	}

	return b, nil
}
