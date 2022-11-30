package mock

import (
	"time"

	"github.com/mgeale/homeserver/internal/db"
)

var mockBalance = &db.Balance{
	ID:          1,
	Name:        "BAL-0022",
	Balance:     100.01,
	BalanceAUD:  100.12,
	PricebookID: 2,
	ProductID:   3,
	Created:     time.Now(),
}

type BalanceModel struct{}

func (m *BalanceModel) Insert(name string, balance, balanceaud float32, pricebook, product int) (int, error) {
	return 2, nil
}

func (m *BalanceModel) Update(id int, name string, balance, balanceaud float32, pricebook, product int) error {
	switch id {
	case 1:
		return nil
	default:
		return db.ErrRecordNotFound
	}
}

func (m *BalanceModel) Get(id int) (*db.Balance, error) {
	switch id {
	case 1:
		return mockBalance, nil
	default:
		return nil, db.ErrRecordNotFound
	}
}

func (m *BalanceModel) Delete(id int) error {
	return nil
}

func (m *BalanceModel) Latest() ([]*db.Balance, error) {
	return []*db.Balance{mockBalance}, nil
}
