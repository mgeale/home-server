package mock

import (
	"time"

	"github.com/mgeale/homeserver/graph/model"
	"github.com/mgeale/homeserver/internal/db"
)

var mockBalances = []*db.Balance{
	{
		ID:          "1",
		Name:        "BAL-0022",
		Balance:     100.01,
		BalanceAUD:  100.12,
		PricebookID: "2",
		ProductID:   "3",
		Created:     time.Now(),
	},
	{
		ID:          "7a59f3e8-b0b9-11ed-a356-0242ac110002",
		Name:        "BAL-0001",
		Balance:     200.00,
		BalanceAUD:  400.00,
		PricebookID: "2222",
		ProductID:   "3333",
		Created:     time.Date(2000, 2, 1, 12, 30, 0, 0, time.UTC),
	},
	{
		ID:          "7a59f5c1-b0b9-11ed-a356-0242ac110002",
		Name:        "BAL-0002",
		Balance:     1000.10,
		BalanceAUD:  2000.20,
		PricebookID: "2222",
		ProductID:   "3333",
		Created:     time.Date(2000, 2, 1, 12, 30, 0, 0, time.UTC),
	},
}

type BalanceModel struct{}

func (m *BalanceModel) Insert(input []*model.InsertBalance) ([]string, error) {
	return []string{"5555", "c9d7bc84-bfeb-11ed-a4d4-0242ac110002"}, nil
}

func (m *BalanceModel) Update(input []*model.UpdateBalance) error {
	switch input[0].ExternalID {
	case "5555":
		return nil
	default:
		return db.ErrRecordNotFound
	}
}

func (m *BalanceModel) Get(query *db.Query) ([]*db.Balance, error) {
	return mockBalances, nil
}

func (m *BalanceModel) Delete(ids []string) error {
	return nil
}
