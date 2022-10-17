package mock

import (
	"time"

	"github.com/mgeale/homeserver/pkg/models"
)

var mockBalance = &models.Balance{
	ID:          1,
	Name:        "BAL-0022",
	Balance:     100,
	BalanceAUD:  100,
	PricebookID: 2,
	ProductID:   3,
	Created:     time.Now(),
}

type BalanceModel struct{}

func (m *BalanceModel) Insert(name string, balance, balanceaud, pricebook, product int) (int, error) {
	return 2, nil
}

func (m *BalanceModel) Get(id int) (*models.Balance, error) {
	switch id {
	case 1:
		return mockBalance, nil
	default:
		return nil, models.ErrNoRecord
	}
}

func (m *BalanceModel) Latest() ([]*models.Balance, error) {
	return []*models.Balance{mockBalance}, nil
}
