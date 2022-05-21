package mysql

import (
	"database/sql"
	"errors"

	"github.com/mgeale/homeserver/pkg/models"
)

type BalanceModel struct {
	DB *sql.DB
}

func (m *BalanceModel) Insert(title string, balance, balanceaud, pricebook, product int) (int, error) {
	stmt := `INSERT INTO balances (title, balance, balanceaud, pricebook, product, created)
    VALUES(?, ?, ?, ?, ?, UTC_TIMESTAMP()))`

	result, err := m.DB.Exec(stmt, title, balance, balanceaud, pricebook, product)
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
	stmt := `SELECT id, title, balance, balanceaud, pricebook, product, created FROM balances
    WHERE id = ?`

	row := m.DB.QueryRow(stmt, id)

	b := &models.Balance{}

	err := row.Scan(&b.ID, &b.Title, &b.Balance, &b.BalanceAud, &b.PriceBook, &b.Product, &b.Created)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, models.ErrNoRecord
		} else {
			return nil, err
		}
	}

	return b, nil
}

func (m *BalanceModel) Latest() ([]*models.Balance, error) {
	stmt := `SELECT id, title, balance, balanceaud, pricebook, product, created FROM balances
    ORDER BY created DESC LIMIT 10`

	rows, err := m.DB.Query(stmt)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	balances := []*models.Balance{}

	for rows.Next() {
		s := &models.Balance{}
		err = rows.Scan(&s.ID, &s.Title, &s.PriceBook, &s.Product, &s.Created)
		if err != nil {
			return nil, err
		}
		balances = append(balances, s)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return balances, nil
}
