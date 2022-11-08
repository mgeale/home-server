package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/mgeale/homeserver/internal/db"
)

func ping(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("OK"))
}

func TestPing(t *testing.T) {
	// app := newTestApplication(t)
	ts := newTestServer(t, http.HandlerFunc(ping))
	defer ts.Close()

	code, _, body := ts.request(t, "GET", "/ping", nil)

	if code != http.StatusOK {
		t.Errorf("want %d; got %d", http.StatusOK, code)
	}

	if string(body) != "OK" {
		t.Errorf("want body to equal %q", "OK")
	}
}

func TestShowBalance(t *testing.T) {
	// app := newTestApplication(t)
	ts := newTestServer(t, http.HandlerFunc(ping))
	defer ts.Close()

	tests := []struct {
		name     string
		urlPath  string
		wantCode int
		wantBody []byte
	}{
		{"Valid ID", "/balance/1", http.StatusOK, []byte("BAL-0022")},
		{"Non-existent ID", "/balance/2", http.StatusNotFound, nil},
		{"Negative ID", "/balance/-1", http.StatusNotFound, nil},
		{"Decimal ID", "/balance/1.23", http.StatusNotFound, nil},
		{"String ID", "/balance/foo", http.StatusNotFound, nil},
		{"Empty ID", "/balance/", http.StatusNotFound, nil},
		{"Trailing slash", "/balance/1/", http.StatusNotFound, nil},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			code, _, body := ts.request(t, "GET", tt.urlPath, nil)

			if code != tt.wantCode {
				t.Errorf("want %d; got %d", tt.wantCode, code)
			}

			if !bytes.Contains(body, tt.wantBody) {
				t.Errorf("want body %s to contain %q", body, tt.wantBody)
			}
		})
	}
}

func TestUpdateBalance(t *testing.T) {
	// app := newTestApplication(t)

	ts := newTestServer(t, http.HandlerFunc(ping))
	defer ts.Close()

	bal := &db.Balance{
		Name:        "BAL-0022",
		Balance:     100,
		BalanceAUD:  1000,
		PricebookID: 3333,
		ProductID:   2222,
	}

	trans := &db.Transaction{
		Name:   "name",
		Amount: 100,
		Date:   "2018-12-23 17:25:22",
		Type:   "Repayment",
	}

	tests := []struct {
		name     string
		urlPath  string
		body     any
		wantCode int
		wantBody []byte
	}{
		{"Valid ID - bal", "/balance/1", bal, http.StatusNoContent, nil},
		{"Non-existent ID - bal", "/balance/2", bal, http.StatusNotFound, nil},
		{"Missing body - bal", "/balance/1", nil, http.StatusBadRequest, nil},
		{"Incomplete body - bal", "/balance/1", &db.Balance{
			Name:       "BAL-0022",
			Balance:    100,
			BalanceAUD: 1000,
		}, http.StatusBadRequest, nil},
		{"Valid ID", "/transaction/1", trans, http.StatusNoContent, nil},
		{"Non-existent ID", "/transaction/2", trans, http.StatusNotFound, nil},
		{"Missing body", "/transaction/1", nil, http.StatusBadRequest, nil},
		{"Incomplete body", "/transaction/1", &db.Transaction{
			Name:   "name",
			Amount: 100,
			Type:   "Repayment",
		}, http.StatusBadRequest, nil},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data, _ := json.Marshal(tt.body)
			r := bytes.NewReader(data)
			code, _, body := ts.request(t, "PUT", tt.urlPath, r)

			if code != tt.wantCode {
				t.Errorf("want %d; got %d", tt.wantCode, code)
			}

			if !bytes.Contains(body, tt.wantBody) {
				t.Errorf("want body %s to contain %q", body, tt.wantBody)
			}
		})
	}
}
