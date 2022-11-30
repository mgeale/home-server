package db

import (
	"log"
	"os"
	"testing"
)

func TestPermissionModelGetAllForUser(t *testing.T) {
	if testing.Short() {
		t.Skip("mysql: skipping integration test")
	}

	tests := []struct {
		name            string
		userID          int64
		wantPermissions int
	}{
		{
			name:            "Valid ID",
			userID:          1,
			wantPermissions: 1,
		},
		{
			name:            "Non-existent ID",
			userID:          22222222,
			wantPermissions: 0,
		},
	}

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, teardown := newTestDB(t)
			defer teardown()

			m := PermissionModel{db, infoLog, errorLog}

			permissions, _ := m.GetAllForUser(tt.userID)
			// TODO: add test for error

			len := len(permissions)

			if len != tt.wantPermissions {
				t.Errorf("want %v; got %v", tt.wantPermissions, len)
			}
		})
	}
}

func TestPermissionModelAddForUser(t *testing.T) {
	if testing.Short() {
		t.Skip("mysql: skipping integration test")
	}

	tests := []struct {
		name      string
		userID    int64
		wantError error
	}{
		{
			name:      "Valid ID",
			userID:    1,
			wantError: nil,
		},
	}

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, teardown := newTestDB(t)
			defer teardown()

			m := PermissionModel{db, infoLog, errorLog}

			err := m.AddForUser(tt.userID, "balances:write")
			// TODO: add test for not found user?

			if err != tt.wantError {
				t.Errorf("want %v; got %s", tt.wantError, err)
			}
		})
	}
}
