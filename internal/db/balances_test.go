package db

import (
	"log"
	"os"
	"testing"

	"github.com/mgeale/homeserver/graph/model"
)

func TestBalanceModelInsert(t *testing.T) {
	if testing.Short() {
		t.Skip("mysql: skipping integration test")
	}

	tests := []struct {
		name      string
		balance   *model.NewBalance
		wantError error
	}{
		{
			name: "Valid ID",
			balance: &model.NewBalance{
				Name:        "BAL-7878",
				Balance:     45.13,
				Balanceaud:  1056.72,
				Pricebookid: "01s9D000001lX8rQAE",
				Productid:   "01t9D000003rsQoQAI",
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

			_, err := m.Insert(tt.balance)

			if err != tt.wantError {
				t.Errorf("want %v; got %s", tt.wantError, err)
			}
		})
	}
}

func TestBalanceModelGet(t *testing.T) {
	if testing.Short() {
		t.Skip("mysql: skipping integration test")
	}

	tests := []struct {
		name      string
		wantError error
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

			m := BalanceModel{db, infoLog, errorLog}

			query := &Query{
				Filters: &Filter{
					Subfilters: []*Filter{
						{
							Subfilters: []*Filter{
								{
									Field: Field("balance"),
									Kind:  Equal,
									Value: 100.89,
								},
								{
									Field: Field("balanceaud"),
									Kind:  Equal,
									Value: 1000.01,
								},
							},
							Kind: And,
						},
						{
							Field: Field("id"),
							Kind:  Equal,
							Value: "7a59f5c1-b0b9-11ed-a356-0242ac110002",
						},
					},
					Kind: Or,
				},
				Sort: Sort{
					Field:     Field("created"),
					Direction: Ascending,
				},
			}

			balances, err := m.Get(query)

			if err != tt.wantError {
				t.Errorf("want %v; got %s", tt.wantError, err)
			}

			if len(balances) != 2 {
				t.Errorf("want 2; got %v", len(balances))
			}

			// if !reflect.DeepEqual(balances[0], tt.wantBalance) {
			// 	t.Errorf("want %v; got %v", tt.wantBalance, balances[0])
			// }
		})
	}
}

func TestBalanceModelUpdate(t *testing.T) {
	if testing.Short() {
		t.Skip("mysql: skipping integration test")
	}

	tests := []struct {
		name      string
		balanceID string
		balance   map[string]interface{}
		wantError error
	}{
		{
			name:      "Valid ID",
			balanceID: "7a59f3e8-b0b9-11ed-a356-0242ac110002",
			balance: map[string]interface{}{
				"name":       "BAL-0022",
				"balance":    200,
				"balanceaud": 2000,
			},
			wantError: nil,
		},
		{
			name:      "Non-existent ID",
			balanceID: "99999999",
			balance: map[string]interface{}{
				"name": "BAL-0022",
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

			m := BalanceModel{db, infoLog, errorLog}

			err := m.Update(tt.balanceID, tt.balance)

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
		balanceID string
		wantError error
	}{
		{
			name:      "Valid ID",
			balanceID: "7a59f3e8-b0b9-11ed-a356-0242ac110002",
			wantError: nil,
		},
		{
			name:      "Non-existent ID",
			balanceID: "9999999",
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
