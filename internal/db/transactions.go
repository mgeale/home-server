package db

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/ido50/sqlz"
)

type Transaction struct {
	ID      string
	Name    string
	Amount  float64
	Date    string
	Type    string
	Created time.Time
}

//TODO: create Type enum

type TransactionModel struct {
	DB       *sql.DB
	InfoLog  *log.Logger
	ErrorLog *log.Logger
}

func (m *TransactionModel) Insert(name string, amount float64, date string, transactionType string) (string, error) {
	stmt := sqlz.New(m.DB, "mysql").
		InsertInto("transactions").
		Columns("id", "name", "amount", "date", "type", "created").
		Values(sqlz.Indirect("UUID()"), name, amount, date, transactionType, sqlz.Indirect("UTC_TIMESTAMP()"))

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

func (m *TransactionModel) Update(id, name string, amount float64, date string, transactionType string) error {
	stmt := sqlz.New(m.DB, "mysql").
		Update("transactions").
		Set("name", name).
		Set("amount", amount).
		Set("date", date).
		Set("type", transactionType).
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

func (m *TransactionModel) Get(query *Query) ([]*Transaction, error) {
	if err := checkOrderQuery(query); err != nil {
		return nil, err
	}

	stmt, err := addOptsToSelectOrdersQuery(sqlz.New(m.DB, "mysql").Select("*").From("transactions"), query)
	if err != nil {
		return nil, err
	}

	var rows []*Transaction
	err = stmt.GetAll(&rows)
	if err != nil {
		return nil, err
	}

	return rows, nil
}

func (m *TransactionModel) Delete(id string) error {
	stmt := sqlz.New(m.DB, "mysql").
		DeleteFrom("transactions").
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
