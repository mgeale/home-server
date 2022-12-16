package db

import (
	"database/sql"
	"errors"
	"log"
	"time"
)

type Transaction struct {
	ID      int
	Name    string
	Amount  float64
	Date    string
	Type    string
	Created time.Time
}

type TransactionModel struct {
	DB       *sql.DB
	InfoLog  *log.Logger
	ErrorLog *log.Logger
}

func (m *TransactionModel) Insert(name string, amount float64, date string, transactionType string) (int, error) {
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

func (m *TransactionModel) Update(id int, name string, amount float64, date string, transactionType string) error {
	stmt := `UPDATE transactions SET name = ?, amount = ?, date = ?, type = ? WHERE id = ?`

	result, err := m.DB.Exec(stmt, name, amount, date, transactionType, id)
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

func (m *TransactionModel) Get(id int) (*Transaction, error) {
	stmt := `SELECT id, name, amount, date, type, created FROM transactions
    WHERE id = ?`

	row := m.DB.QueryRow(stmt, id)

	t := &Transaction{}

	err := row.Scan(&t.ID, &t.Name, &t.Amount, &t.Date, &t.Type, &t.Created)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrRecordNotFound
		} else {
			return nil, err
		}
	}

	return t, nil
}

func (m *TransactionModel) Delete(id int) error {
	stmt := `DELETE FROM transactions WHERE id = ?`

	result, err := m.DB.Exec(stmt, id)
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

func (m *TransactionModel) Latest() ([]*Transaction, error) {
	stmt := `SELECT id, name, amount, date, type, created FROM transactions
    ORDER BY created DESC LIMIT 10`

	rows, err := m.DB.Query(stmt)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	ts := []*Transaction{}

	for rows.Next() {
		s := &Transaction{}
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
