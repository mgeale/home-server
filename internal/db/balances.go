package db

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/ido50/sqlz"
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

func (m *BalanceModel) Insert(name string, balance, balanceaud float64, pricebookid, productid string) (string, error) {
	stmt := sqlz.New(m.DB, "mysql").
		InsertInto("balances").
		Columns("id", "name", "balance", "balanceaud", "pricebookid", "productid", "created").
		Values(sqlz.Indirect("UUID()"), name, balance, balanceaud, pricebookid, productid, sqlz.Indirect("UTC_TIMESTAMP()"))

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

func (m *BalanceModel) Update(id, name string, balance, balanceaud float64, pricebookid, productid string) error {
	stmt := sqlz.New(m.DB, "mysql").
		Update("balances").
		Set("name", name).
		Set("balance", balance).
		Set("balanceaud", balanceaud).
		Set("pricebookid", pricebookid).
		Set("productid", productid).
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
