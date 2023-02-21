package db

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/ido50/sqlz"
	"github.com/mgeale/homeserver/graph/model"
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

func (m *TransactionModel) Insert(input *model.NewTransaction) (string, error) {
	stmt := sqlz.New(m.DB, "mysql").
		InsertInto("transactions").
		Columns("id", "name", "amount", "date", "type", "created").
		Values(sqlz.Indirect("UUID()"), input.Name, input.Amount, input.Date, input.Type, sqlz.Indirect("UTC_TIMESTAMP()"))

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

func (m *TransactionModel) Update(id string, values map[string]interface{}) error {
	stmt := sqlz.New(m.DB, "mysql").
		Update("transactions").
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
