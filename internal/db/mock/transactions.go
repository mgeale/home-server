package mock

import (
	"time"

	"github.com/mgeale/homeserver/graph/model"
	"github.com/mgeale/homeserver/internal/db"
)

var mockTransaction = &db.Transaction{
	ID:      "1",
	Name:    "BAL-0022",
	Amount:  100,
	Date:    "2018-12-23 17:25:22",
	Type:    "Repayment",
	Created: time.Now(),
}

type TransactionModel struct{}

func (m *TransactionModel) Insert(input *model.NewTransaction) (string, error) {
	return "2", nil
}

func (m *TransactionModel) Update(id string, values map[string]interface{}) error {
	switch id {
	case "1":
		return nil
	default:
		return db.ErrRecordNotFound
	}
}

func (m *TransactionModel) Get(query *db.Query) ([]*db.Transaction, error) {
	return []*db.Transaction{mockTransaction}, nil
}

func (m *TransactionModel) Delete(id string) error {
	return nil
}
