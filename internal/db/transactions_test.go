package db

import (
	"log"
	"os"
	"testing"

	"github.com/mgeale/homeserver/graph/model"
)

func TestTransactionModelGet(t *testing.T) {
	if testing.Short() {
		t.Skip("mysql: skipping integration test")
	}

	tests := []struct {
		name            string
		wantTransaction *Transaction
		wantError       error
	}{
		{
			name:      "Valid ID",
			wantError: nil,
		},
	}

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, teardown := newTestDB(t)
			defer teardown()

			m := TransactionModel{db, infoLog, errorLog}

			query := &Query{
				Filters: &Filter{
					Field: Field("Amount"),
					Kind:  Equal,
					Value: 100.00,
				},
				Sort: Sort{
					Field:     Field("created"),
					Direction: Ascending,
				},
				Limit: 1,
			}

			transactions, err := m.Get(query)

			if err != tt.wantError {
				t.Errorf("want %v; got %s", tt.wantError, err)
			}

			if len(transactions) != 1 {
				t.Errorf("want at only 1 balance")
			}
		})
	}
}

func TestTransactionModelInsert(t *testing.T) {
	if testing.Short() {
		t.Skip("mysql: skipping integration test")
	}

	tests := []struct {
		name        string
		transaction *model.NewTransaction
		wantError   error
	}{
		{
			name: "Valid Transaction",
			transaction: &model.NewTransaction{
				Name:   "TNS-00999",
				Amount: 111.00,
				Date:   "2018-12-23 17:25:22",
				Type:   "Repayment",
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

			m := TransactionModel{db, infoLog, errorLog}

			_, err := m.Insert(tt.transaction)

			if err != tt.wantError {
				t.Errorf("want %v; got %s", tt.wantError, err)
			}
		})
	}
}

func TestTransactionModelUpdate(t *testing.T) {
	if testing.Short() {
		t.Skip("mysql: skipping integration test")
	}

	tests := []struct {
		name        string
		transID     string
		transaction map[string]interface{}
		wantError   error
	}{
		{
			name:    "Valid ID",
			transID: "1c0d2b44-b0ce-11ed-b95f-dca632bb7cae",
			transaction: map[string]interface{}{
				"name":   "TNS-00999",
				"amount": 111.00,
				"date":   "2018-12-23 17:25:22",
				"type":   "Repayment",
			},
			wantError: nil,
		},
		{
			name:    "Non-existent ID",
			transID: "99999999",
			transaction: map[string]interface{}{
				"name":   "TNS-00999",
				"amount": 111.00,
				"date":   "2018-12-23 17:25:22",
				"type":   "Repayment",
			},
			wantError: ErrRecordNotFound,
		},
	}

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, teardown := newTestDB(t)
			defer teardown()

			m := TransactionModel{db, infoLog, errorLog}

			err := m.Update(tt.transID, tt.transaction)

			if err != tt.wantError {
				t.Errorf("want %v; got %s", tt.wantError, err)
			}
		})
	}
}

func TestTransactionModelDelete(t *testing.T) {
	if testing.Short() {
		t.Skip("mysql: skipping integration test")
	}

	tests := []struct {
		name      string
		transID   string
		wantError error
	}{
		{
			name:      "Valid ID",
			transID:   "1c0d2b44-b0ce-11ed-b95f-dca632bb7cae",
			wantError: nil,
		},
		{
			name:      "Non-existent ID",
			transID:   "22222222",
			wantError: ErrRecordNotFound,
		},
	}

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, teardown := newTestDB(t)
			defer teardown()

			m := TransactionModel{db, infoLog, errorLog}

			err := m.Delete(tt.transID)

			if err != tt.wantError {
				t.Errorf("want %v; got %s", tt.wantError, err)
			}
		})
	}
}
