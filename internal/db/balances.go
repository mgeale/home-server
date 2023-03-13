package db

import (
	"database/sql"
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

func (m *BalanceModel) Insert(input []*model.InsertBalance) ([]string, error) {
	stmt := sqlz.New(m.DB, "mysql").
		InsertInto("balances").
		Columns("id", "name", "balance", "balanceaud", "pricebookid", "productid", "created").
		Returning("id")

	vals := make([][]interface{}, len(input))
	for i, value := range input {
		vals[i] = []interface{}{sqlz.Indirect("UUID()"), value.Name, value.Balance, value.Balanceaud, value.Pricebookid, value.Productid, sqlz.Indirect("UTC_TIMESTAMP()")}
	}
	stmt.ValueMultiple(vals)

	var ids []string
	err := stmt.GetAll(&ids)
	if err != nil {
		return nil, err
	}

	return ids, nil
}

func (m *BalanceModel) Update(input []*model.UpdateBalance) error {
	return sqlz.New(m.DB, "mysql").Transactional(func(tx *sqlz.Tx) error {
		for _, in := range input {
			values := constructValuesMap(*in)
			delete(values, "externalid")
			delete(values, "displayurl")

			stmt := tx.Update("balances").
				SetMap(values).
				Where(sqlz.Eq("id", in.ExternalID))

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
		}

		return nil
	})
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

func (m *BalanceModel) Delete(ids []string) error {
	id := make([]interface{}, len(ids))
	for i, v := range ids {
		id[i] = v
	}
	stmt := sqlz.New(m.DB, "mysql").
		DeleteFrom("balances").
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
