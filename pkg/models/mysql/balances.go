package mysql

import (
	"database/sql"
	"errors"

	"github.com/mgeale/homeserver/pkg/models"
)

type BalanceModel struct {
	DB *sql.DB
}

func (m *BalanceModel) Insert(name string, balance, balanceaud, pricebookid, productid int) (int, error) {
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

func (m *BalanceModel) Get(id int) (*models.Balance, error) {
	stmt := `SELECT id, name, balance, balanceaud, pricebookid, productid, created FROM balances
    WHERE id = ?`

	row := m.DB.QueryRow(stmt, id)

	s := &models.Balance{}

	err := row.Scan(&s.ID, &s.Name, &s.Balance, &s.BalanceAUD, &s.PricebookID, &s.ProductID, &s.Created)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, models.ErrNoRecord
		} else {
			return nil, err
		}
	}

	return s, nil
}

func (m *BalanceModel) Latest() ([]*models.Balance, error) {
	stmt := `SELECT id, name, balance, balanceaud, pricebookid, productid, created FROM balances
    ORDER BY created DESC LIMIT 10`

	rows, err := m.DB.Query(stmt)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	snippets := []*models.Balance{}

	for rows.Next() {
		s := &models.Balance{}
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
