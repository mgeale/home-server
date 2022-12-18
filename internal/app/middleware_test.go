package app

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/mgeale/homeserver/internal/db"
	"github.com/mgeale/homeserver/internal/db/mock"
	"github.com/mgeale/homeserver/internal/jsonlog"
)

func TestBasicAuth(t *testing.T) {
	if testing.Short() {
		t.Skip("mysql: skipping integration test")
	}

	tests := []struct {
		name           string
		userEmail      string
		userPassword   string
		wantStatusCode int
	}{
		{
			name:           "Valid User",
			userEmail:      "alice@example.com",
			userPassword:   "password",
			wantStatusCode: http.StatusOK,
		},
		{
			name:           "Valid User",
			userEmail:      "alice@example.com",
			userPassword:   "wrongpassword",
			wantStatusCode: http.StatusUnauthorized,
		},
		{
			name:           "Not-found User",
			userEmail:      "doesntexist@example.com",
			userPassword:   "password",
			wantStatusCode: http.StatusNotFound,
		},
		{
			name:           "Non-valid User",
			wantStatusCode: http.StatusUnauthorized,
		},
	}

	app := &Application{
		Logger: jsonlog.NewLogger(os.Stdout, jsonlog.LevelInfo),
		Config: Config{},
		Models: db.Models{
			Balances:     &mock.BalanceModel{},
			Transactions: &mock.TransactionModel{},
			Users:        &mock.UserModel{},
			Permissions:  &mock.PermissionModel{},
		},
	}

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user := app.contextGetUser(r)
		w.Write([]byte(user.Email))
	})

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rr := httptest.NewRecorder()

			r, err := http.NewRequest(http.MethodGet, "/", nil)
			if err != nil {
				t.Fatal(err)
			}

			if tt.userEmail != "" {
				r.SetBasicAuth(tt.userEmail, tt.userPassword)
			}

			app.BasicAuth(handler).ServeHTTP(rr, r)
			rs := rr.Result()

			if rs.StatusCode != tt.wantStatusCode {
				t.Errorf("want %d; got %d", tt.wantStatusCode, rs.StatusCode)
			}
		})
	}
}
