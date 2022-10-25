package mysql

import (
	"reflect"
	"testing"
	"time"

	"github.com/mgeale/homeserver/pkg/models"
)

func TestBalanceModelGet(t *testing.T) {
	if testing.Short() {
		t.Skip("mysql: skipping integration test")
	}

	tests := []struct {
		name        string
		balanceID   int
		wantBalance *models.Balance
		wantError   error
	}{
		{
			name:      "Valid ID",
			balanceID: 1,
			wantBalance: &models.Balance{
				ID:          1,
				Name:        "BAL-0022",
				Balance:     100.89,
				BalanceAUD:  1000.01,
				PricebookID: 3333,
				ProductID:   2222,
				Created:     time.Date(2018, 12, 23, 17, 25, 22, 0, time.UTC),
			},
			wantError: nil,
		},
		{
			name:        "Zero ID",
			balanceID:   0,
			wantBalance: nil,
			wantError:   models.ErrNoRecord,
		},
		{
			name:        "Non-existent ID",
			balanceID:   2,
			wantBalance: nil,
			wantError:   models.ErrNoRecord,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, teardown := newTestDB(t)
			defer teardown()

			m := BalanceModel{db}

			balance, err := m.Get(tt.balanceID)

			if err != tt.wantError {
				t.Errorf("want %v; got %s", tt.wantError, err)
			}

			if !reflect.DeepEqual(balance, tt.wantBalance) {
				t.Errorf("want %v; got %v", tt.wantBalance, balance)
			}
		})
	}
}

func TestBalanceModelUpdate(t *testing.T) {
	if testing.Short() {
		t.Skip("mysql: skipping integration test")
	}

	tests := []struct {
		name      string
		balanceID int
		wantError error
	}{
		{
			name:      "Valid ID",
			balanceID: 1,
			wantError: nil,
		},
		{
			name:      "Non-existent ID",
			balanceID: 2,
			wantError: models.ErrNoRecord,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, teardown := newTestDB(t)
			defer teardown()

			m := BalanceModel{db}

			err := m.Update(tt.balanceID, "BAL-0022", 200, 2000, 3333, 2222)

			if err != tt.wantError {
				t.Errorf("want %v; got %s", tt.wantError, err)
			}
		})
	}
}

func TestBalanceModelDelete(t *testing.T) {
	if testing.Short() {
		t.Skip("mysql: skipping integration test")
	}

	tests := []struct {
		name      string
		balanceID int
		wantError error
	}{
		{
			name:      "Valid ID",
			balanceID: 1,
			wantError: nil,
		},
		{
			name:      "Non-existent ID",
			balanceID: 2,
			wantError: models.ErrNoRecord,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, teardown := newTestDB(t)
			defer teardown()

			m := BalanceModel{db}

			err := m.Delete(tt.balanceID)

			if err != tt.wantError {
				t.Errorf("want %v; got %s", tt.wantError, err)
			}
		})
	}
}
