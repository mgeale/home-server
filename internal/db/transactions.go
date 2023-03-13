package db

import (
	"database/sql"
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

func (m *TransactionModel) Insert(input []*model.InsertTransaction) ([]string, error) {
	stmt := sqlz.New(m.DB, "mysql").
		InsertInto("transactions").
		Columns("id", "name", "amount", "date", "type", "created").
		Returning("id")

	vals := make([][]interface{}, len(input))
	for i, value := range input {
		vals[i] = []interface{}{sqlz.Indirect("UUID()"), value.Name, value.Amount, value.Date, value.Type, sqlz.Indirect("UTC_TIMESTAMP()")}
	}
	stmt.ValueMultiple(vals)

	var ids []string
	err := stmt.GetAll(&ids)
	if err != nil {
		return nil, err
	}

	return ids, nil
}

func (m *TransactionModel) Update(input []*model.UpdateTransaction) error {
	return sqlz.New(m.DB, "mysql").Transactional(func(tx *sqlz.Tx) error {
		for _, in := range input {
			values := constructValuesMap(*in)
			delete(values, "externalid")
			delete(values, "displayurl")

			result, err := tx.Update("transactions").
				SetMap(values).
				Where(sqlz.Eq("id", in.ExternalID)).
				Exec()

			if err != nil {
				return err
			}

			n, err := result.RowsAffected()
			if n == 0 {
				return ErrRecordNotFound
			} else if err != nil {
				return err
			}
		}

		return nil
	})
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

func (m *TransactionModel) Delete(ids []string) error {
	id := make([]interface{}, len(ids))
	for i, v := range ids {
		id[i] = v
	}

	stmt := sqlz.New(m.DB, "mysql").
		DeleteFrom("transactions").
		Where(sqlz.In("id", id...))

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
