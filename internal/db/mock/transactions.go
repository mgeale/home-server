package mock

import (
	"time"

	"github.com/mgeale/homeserver/graph/model"
	"github.com/mgeale/homeserver/internal/db"
)

var mockTransactions = []*db.Transaction{
	{
		ID:      "1",
		Name:    "TRANS-0001",
		Amount:  100,
		Date:    "2018-12-23 17:25:22",
		Type:    "Repayment",
		Created: time.Now(),
	},
	{
		ID:      "2",
		Name:    "TRANS-0002",
		Amount:  200.00,
		Date:    "2018-12-23 17:25:22",
		Type:    "Repayment",
		Created: time.Date(2000, 2, 1, 12, 30, 0, 0, time.UTC),
	},
	{
		ID:      "3",
		Name:    "TRANS-0003",
		Amount:  400.00,
		Date:    "2018-12-23 17:25:22",
		Type:    "Repayment",
		Created: time.Date(2000, 2, 1, 12, 30, 0, 0, time.UTC),
	},
}

type TransactionModel struct{}

func (m *TransactionModel) Insert(input []*model.InsertTransaction) ([]string, error) {
	return []string{"2"}, nil
}

func (m *TransactionModel) Update(input []*model.UpdateTransaction) error {
	switch input[0].ExternalID {
	case "5555":
		return nil
	default:
		return db.ErrRecordNotFound
	}
}

func (m *TransactionModel) Get(query *db.Query) ([]*db.Transaction, error) {
	return mockTransactions, nil
}

func (m *TransactionModel) Delete(ids []string) error {
	return nil
}
