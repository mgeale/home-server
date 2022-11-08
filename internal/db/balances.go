package db

import (
	"database/sql"
	"errors"
	"log"
	"time"
)

type Balance struct {
	ID          int
	Name        string
	Balance     float64
	BalanceAUD  float64
	PricebookID int
	ProductID   int
	Created     time.Time
}

type BalanceModel struct {
	DB       *sql.DB
	InfoLog  *log.Logger
	ErrorLog *log.Logger
}

func (m *BalanceModel) Insert(name string, balance, balanceaud float64, pricebookid, productid int) (int, error) {
	stmt := `INSERT INTO balances (name, balance, balanceaud, pricebookid, productid, created)
    VALUES(?, ?, ?, ?, ?, UTC_TIMESTAMP())`

	result, err := m.DB.Exec(stmt, name, balance, balanceaud, pricebookid, productid)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (m *BalanceModel) Update(id int, name string, balance, balanceaud float64, pricebookid, productid int) error {
	stmt := `UPDATE balances SET name = ?, balance = ?, balanceaud = ?, pricebookid = ?, productid = ? WHERE id = ?`

	result, err := m.DB.Exec(stmt, name, balance, balanceaud, pricebookid, productid, id)
	if err != nil {
		return err
	}

	n, err := result.RowsAffected()
	if n == 0 {
		return ErrNoRecord
	} else if err != nil {
		return err
	}

	return nil
}

func (m *BalanceModel) Get(id int) (*Balance, error) {
	stmt := `SELECT id, name, balance, balanceaud, pricebookid, productid, created FROM balances
    WHERE id = ?`

	row := m.DB.QueryRow(stmt, id)

	b := &Balance{}

	err := row.Scan(&b.ID, &b.Name, &b.Balance, &b.BalanceAUD, &b.PricebookID, &b.ProductID, &b.Created)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNoRecord
		} else {
			return nil, err
		}
	}

	return b, nil
}

func (m *BalanceModel) Delete(id int) error {
	stmt := `DELETE FROM balances WHERE id = ?`

	result, err := m.DB.Exec(stmt, id)
	if err != nil {
		return err
	}

	n, err := result.RowsAffected()
	if n == 0 {
		return ErrNoRecord
	} else if err != nil {
		return err
	}

	return nil
}

func (m *BalanceModel) Latest() ([]*Balance, error) {
	stmt := `SELECT id, name, balance, balanceaud, pricebookid, productid, created FROM balances
    ORDER BY created DESC LIMIT 10`

	rows, err := m.DB.Query(stmt)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	snippets := []*Balance{}

	for rows.Next() {
		s := &Balance{}
		err = rows.Scan(&s.ID, &s.Name, &s.Balance, &s.BalanceAUD, &s.PricebookID, &s.ProductID, &s.Created)
		if err != nil {
			return nil, err
		}
		snippets = append(snippets, s)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return snippets, nil
}
