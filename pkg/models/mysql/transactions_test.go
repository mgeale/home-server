package mysql

import (
	"reflect"
	"testing"
	"time"

	"github.com/mgeale/homeserver/pkg/models"
)

func TestTransactionModelGet(t *testing.T) {
	if testing.Short() {
		t.Skip("mysql: skipping integration test")
	}

	tests := []struct {
		name            string
		transactionID   int
		wantTransaction *models.Transaction
		wantError       error
	}{
		{
			name:          "Valid ID",
			transactionID: 1,
			wantTransaction: &models.Transaction{
				ID:      1,
				Name:    "name",
				Amount:  100,
				Date:    "2018-12-23 17:25:22",
				Type:    "Repayment",
				Created: time.Date(2018, 12, 23, 17, 25, 22, 0, time.UTC),
			},
			wantError: nil,
		},
		{
			name:            "Zero ID",
			transactionID:   0,
			wantTransaction: nil,
			wantError:       models.ErrNoRecord,
		},
		{
			name:            "Non-existent ID",
			transactionID:   2,
			wantTransaction: nil,
			wantError:       models.ErrNoRecord,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, teardown := newTestDB(t)
			defer teardown()

			m := TransactionModel{db}

			transaction, err := m.Get(tt.transactionID)

			if err != tt.wantError {
				t.Errorf("want %v; got %s", tt.wantError, err)
			}

			if !reflect.DeepEqual(transaction, tt.wantTransaction) {
				t.Errorf("want %v; got %v", tt.wantTransaction, transaction)
			}
		})
	}
}
