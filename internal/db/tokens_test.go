package db

import (
	"log"
	"os"
	"testing"
	"time"
)

func TestTokenModelNew(t *testing.T) {
	if testing.Short() {
		t.Skip("mysql: skipping integration test")
	}

	tests := []struct {
		name      string
		userID    int64
		ttl       time.Duration
		wantError error
	}{
		{
			name:      "Valid",
			userID:    1111111,
			ttl:       10000,
			wantError: nil,
		},
	}

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, teardown := newTestDB(t)
			defer teardown()

			m := TokenModel{db, infoLog, errorLog}

			token, err := m.New(tt.userID, tt.ttl, ScopeAuthentication)

			if err != tt.wantError {
				t.Errorf("want %v; got %s", tt.wantError, err)
			}

			if token != nil {
				if tt.userID != token.UserID {
					t.Errorf("want %v; got %v", tt.userID, token.UserID)
				}
			}
		})
	}
}
