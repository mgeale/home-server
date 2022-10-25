package mock

import (
	"time"

	"github.com/mgeale/homeserver/pkg/models"
)

var mockBalance = &models.Balance{
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
		return models.ErrNoRecord
	}
}

func (m *BalanceModel) Get(id int) (*models.Balance, error) {
	switch id {
	case 1:
		return mockBalance, nil
	default:
		return nil, models.ErrNoRecord
	}
}

func (m *BalanceModel) Delete(id int) error {
	return nil
}

func (m *BalanceModel) Latest() ([]*models.Balance, error) {
	return []*models.Balance{mockBalance}, nil
}
