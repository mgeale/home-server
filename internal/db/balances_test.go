package db

import (
	"log"
	"os"
	"reflect"
	"testing"
	"time"
)

func TestBalanceModelGetById(t *testing.T) {
	if testing.Short() {
		t.Skip("mysql: skipping integration test")
	}

	tests := []struct {
		name        string
		balanceID   int
		wantBalance *Balance
		wantError   error
	}{
		{
			name:      "Valid ID",
			balanceID: 1,
			wantBalance: &Balance{
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
			wantError:   ErrRecordNotFound,
		},
		{
			name:        "Non-existent ID",
			balanceID:   2,
			wantBalance: nil,
			wantError:   ErrRecordNotFound,
		},
	}

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, teardown := newTestDB(t)
			defer teardown()

			m := BalanceModel{db, infoLog, errorLog}

			balance, err := m.GetById(tt.balanceID)

			if err != tt.wantError {
				t.Errorf("want %v; got %s", tt.wantError, err)
			}

			if !reflect.DeepEqual(balance, tt.wantBalance) {
				t.Errorf("want %v; got %v", tt.wantBalance, balance)
			}
		})
	}
}

func TestBalanceModelGet(t *testing.T) {
	if testing.Short() {
		t.Skip("mysql: skipping integration test")
	}

	tests := []struct {
		name        string
		balanceID   int
		wantBalance *Balance
		wantError   error
	}{
		{
			name:      "Valid ID",
			balanceID: 1,
			wantBalance: &Balance{
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
	}

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, teardown := newTestDB(t)
			defer teardown()

			m := BalanceModel{db, infoLog, errorLog}

			query := &Query{
				Filters: []Filter{
					{
						Field: Field("balance"),
						Kind:  Equal,
						Value: 100.89,
					},
				},
				Sort: Sort{
					Field:     Field("created"),
					Direction: Ascending,
				},
				Limit: 1,
			}

			balances, err := m.Get(query)

			if err != tt.wantError {
				t.Errorf("want %v; got %s", tt.wantError, err)
			}

			if len(balances) != 1 {
				t.Errorf("want at only 1 balance")
			}

			if !reflect.DeepEqual(balances[0], tt.wantBalance) {
				t.Errorf("want %v; got %v", tt.wantBalance, balances[0])
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
			wantError: ErrRecordNotFound,
		},
	}

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, teardown := newTestDB(t)
			defer teardown()

			m := BalanceModel{db, infoLog, errorLog}

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
			wantError: ErrRecordNotFound,
		},
	}

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, teardown := newTestDB(t)
			defer teardown()

			m := BalanceModel{db, infoLog, errorLog}

			err := m.Delete(tt.balanceID)

			if err != tt.wantError {
				t.Errorf("want %v; got %s", tt.wantError, err)
			}
		})
	}
}
