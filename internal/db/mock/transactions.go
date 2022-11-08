package mock

import (
	"time"

	"github.com/mgeale/homeserver/internal/db"
)

var mockTransaction = &db.Transaction{
	ID:      1,
	Name:    "BAL-0022",
	Amount:  100,
	Date:    "2018-12-23 17:25:22",
	Type:    "Repayment",
	Created: time.Now(),
}

type TransactionModel struct{}

func (m *TransactionModel) Insert(name string, amount float32, date, transactionType string) (int, error) {
	return 2, nil
}

func (m *TransactionModel) Update(id int, name string, amount float32, date, transactionType string) error {
	switch id {
	case 1:
		return nil
	default:
		return db.ErrNoRecord
	}
}

func (m *TransactionModel) Get(id int) (*db.Transaction, error) {
	switch id {
	case 1:
		return mockTransaction, nil
	default:
		return nil, db.ErrNoRecord
	}
}

func (m *TransactionModel) Delete(id int) error {
	return nil
}

func (m *TransactionModel) Latest() ([]*db.Transaction, error) {
	return []*db.Transaction{mockTransaction}, nil
}
