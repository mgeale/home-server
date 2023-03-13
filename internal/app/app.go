package app

import (
	"context"
	"reflect"
	"strings"
	"sync"

	"github.com/mgeale/homeserver/graph/model"
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

func (app *Application) CreateBalance(ctx context.Context, input []*model.InsertBalance) ([]string, error) {
	ids, err := app.Models.Balances.Insert(input)
	if err != nil {
		return nil, err
	}

	return ids, nil
}

func (app *Application) CreateTransaction(ctx context.Context, input []*model.InsertTransaction) ([]string, error) {
	ids, err := app.Models.Transactions.Insert(input)
	if err != nil {
		return nil, err
	}

	return ids, nil
}

func (app *Application) UpdateBalance(ctx context.Context, input []*model.UpdateBalance) error {
	err := app.Models.Balances.Update(input)
	if err != nil {
		return err
	}

	return nil
}

func (app *Application) UpdateTransaction(ctx context.Context, input []*model.UpdateTransaction) error {
	err := app.Models.Transactions.Update(input)
	if err != nil {
		return err
	}

	return nil
}

func (app *Application) DeleteBalance(ctx context.Context, ids []string) error {
	err := app.Models.Balances.Delete(ids)
	if err != nil {
		return err
	}
	return nil
}

func (app *Application) DeleteTransaction(ctx context.Context, ids []string) error {
	err := app.Models.Transactions.Delete(ids)
	if err != nil {
		return err
	}
	return nil
}

func (app *Application) GetBalances(ctx context.Context, where *model.BalanceFilter, orderBy model.BalanceSort, limit *int) ([]*model.Balance, error) {
	query := createBalanceQuery(where, orderBy, limit)

	balances, err := app.Models.Balances.Get(query)
	if err != nil {
		return nil, err
	}

	return toBalanceModel(balances), nil
}

func (app *Application) GetTransactions(ctx context.Context, where *model.TransactionFilter, orderBy model.TransactionSort, limit *int) ([]*model.Transaction, error) {
	query := createTransactionQuery(where, orderBy, limit)

	transactions, err := app.Models.Transactions.Get(query)
	if err != nil {
		return nil, err
	}

	return toTransactionModel(transactions), nil
}

func constructValuesMap(structure any) map[string]interface{} {
	values := reflect.ValueOf(structure)
	types := values.Type()
	valuesMap := map[string]interface{}{}
	for i := 0; i < values.NumField(); i++ {
		valuesMap[strings.ToLower(types.Field(i).Name)] = values.Field(i).Interface()
	}
	return valuesMap
}
