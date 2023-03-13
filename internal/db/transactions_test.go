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
		name         string
		transactions []*model.InsertTransaction
		wantError    error
	}{
		{
			name: "Valid Transaction",
			transactions: []*model.InsertTransaction{
				{
					Name:   "TNS-00999",
					Amount: 111.00,
					Date:   "2018-12-23 17:25:22",
					Type:   "Repayment",
				},
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

			ids, err := m.Insert(tt.transactions)

			if err != tt.wantError {
				t.Errorf("want %v; got %s", tt.wantError, err)
			}

			if len(ids) != 1 {
				t.Errorf("want 1; got %v", len(ids))
			}
		})
	}
}

func TestTransactionModelUpdate(t *testing.T) {
	if testing.Short() {
		t.Skip("mysql: skipping integration test")
	}

	var name string = "TNS-00999"
	var amount float64 = 111.00
	var date string = "2018-12-23 17:25:22"
	var transType string = "Repayment"

	tests := []struct {
		name         string
		transactions []*model.UpdateTransaction
		wantError    error
	}{
		{
			name: "Valid ID",
			transactions: []*model.UpdateTransaction{
				{
					ExternalID: "1c0d2b44-b0ce-11ed-b95f-dca632bb7cae",
					Name:       &name,
					Amount:     &amount,
					Date:       &date,
					Type:       &transType,
				},
			},
			wantError: nil,
		},
		{
			name: "Non-existent ID",
			transactions: []*model.UpdateTransaction{
				{
					ExternalID: "99999999",
					Name:       &name,
					Amount:     &amount,
					Date:       &date,
					Type:       &transType,
				},
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

			err := m.Update(tt.transactions)

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
		transIDs  []string
		wantError error
	}{
		{
			name:      "Valid ID",
			transIDs:  []string{"1c0d2b44-b0ce-11ed-b95f-dca632bb7cae"},
			wantError: nil,
		},
		{
			name:      "Non-existent ID",
			transIDs:  []string{"22222222"},
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

			err := m.Delete(tt.transIDs)

			if err != tt.wantError {
				t.Errorf("want %v; got %s", tt.wantError, err)
			}
		})
	}
}
