package mock

import (
	"time"

	"github.com/mgeale/homeserver/graph/model"
	"github.com/mgeale/homeserver/internal/db"
)

var mockBalance = &db.Balance{
	ID:          "1",
	Name:        "BAL-0022",
	Balance:     100.01,
	BalanceAUD:  100.12,
	PricebookID: "2",
	ProductID:   "3",
	Created:     time.Now(),
}

type BalanceModel struct{}

func (m *BalanceModel) Insert(input *model.NewBalance) (string, error) {
	return "2", nil
}

func (m *BalanceModel) Update(id string, values map[string]interface{}) error {
	switch id {
	case "1":
		return nil
	default:
		return db.ErrRecordNotFound
	}
}

func (m *BalanceModel) Get(query *db.Query) ([]*db.Balance, error) {
	return []*db.Balance{mockBalance}, nil
}

func (m *BalanceModel) Delete(id string) error {
	return nil
}
