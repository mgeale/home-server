package db

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/ido50/sqlz"
	"github.com/mgeale/homeserver/graph/model"
)

type Balance struct {
	ID          string
	Name        string
	Balance     float64
	BalanceAUD  float64
	PricebookID string
	ProductID   string
	Created     time.Time
}

type BalanceModel struct {
	DB       *sql.DB
	InfoLog  *log.Logger
	ErrorLog *log.Logger
}

func (m *BalanceModel) Insert(input *model.NewBalance) (string, error) {
	stmt := sqlz.New(m.DB, "mysql").
		InsertInto("balances").
		Columns("id", "name", "balance", "balanceaud", "pricebookid", "productid", "created").
		Values(sqlz.Indirect("UUID()"), input.Name, input.Balance, input.Balanceaud, input.Pricebookid, input.Productid, sqlz.Indirect("UTC_TIMESTAMP()"))

	result, err := stmt.Exec()
	if err != nil {
		return "0", err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return "0", err
	}

	return fmt.Sprint(id), nil
}

func (m *BalanceModel) Update(id string, values map[string]interface{}) error {
	stmt := sqlz.New(m.DB, "mysql").
		Update("balances").
		SetMap(values).
		Where(sqlz.Eq("id", id))

	result, err := stmt.Exec()
	if err != nil {
		return err
	}

	n, err := result.RowsAffected()
	if n == 0 {
		return ErrRecordNotFound
	} else if err != nil {
		return err
	}

	return nil
}

func (m *BalanceModel) Get(query *Query) ([]*Balance, error) {
	if err := checkOrderQuery(query); err != nil {
		return nil, err
	}

	stmt, err := addOptsToSelectOrdersQuery(sqlz.New(m.DB, "mysql").Select("*").From("balances"), query)
	if err != nil {
		return nil, err
	}

	var rows []*Balance
	err = stmt.GetAll(&rows)
	if err != nil {
		return nil, err
	}

	return rows, nil
}

func (m *BalanceModel) Delete(id string) error {
	stmt := sqlz.New(m.DB, "mysql").
		DeleteFrom("balances").
		Where(sqlz.Eq("id", id))

	result, err := stmt.Exec()
	if err != nil {
		return err
	}

	n, err := result.RowsAffected()
	if n == 0 {
		return ErrRecordNotFound
	} else if err != nil {
		return err
	}

	return nil
}
