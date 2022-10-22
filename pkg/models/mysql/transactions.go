package mysql

import (
	"database/sql"
	"errors"

	"github.com/mgeale/homeserver/pkg/models"
)

type TransactionModel struct {
	DB *sql.DB
}

func (m *TransactionModel) Insert(name string, amount int, date string, transactionType string) (int, error) {
	stmt := `INSERT INTO transactions (name, amount, date, type, created)
    VALUES(?, ?, ?, ?, UTC_TIMESTAMP())`

	result, err := m.DB.Exec(stmt, name, amount, date, transactionType)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (m *TransactionModel) Get(id int) (*models.Transaction, error) {
	stmt := `SELECT id, name, amount, date, type, created FROM transactions
    WHERE id = ?`

	row := m.DB.QueryRow(stmt, id)

	t := &models.Transaction{}

	err := row.Scan(&t.ID, &t.Name, &t.Amount, &t.Date, &t.Type, &t.Created)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, models.ErrNoRecord
		} else {
			return nil, err
		}
	}

	return t, nil
}

func (m *TransactionModel) Latest() ([]*models.Transaction, error) {
	stmt := `SELECT id, name, amount, date, type, created FROM transactions
    ORDER BY created DESC LIMIT 10`

	rows, err := m.DB.Query(stmt)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	ts := []*models.Transaction{}

	for rows.Next() {
		s := &models.Transaction{}
		err = rows.Scan(&s.ID, &s.Name, &s.Amount, &s.Date, &s.Type, &s.Created)
		if err != nil {
			return nil, err
		}
		ts = append(ts, s)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return ts, nil
}
