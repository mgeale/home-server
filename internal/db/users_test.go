package db

import (
	"log"
	"os"
	"testing"
)

func TestUserModelGetByEmail(t *testing.T) {
	if testing.Short() {
		t.Skip("mysql: skipping integration test")
	}

	tests := []struct {
		name      string
		userEmail string
		wantUser  *User
		wantError error
	}{
		{
			name:      "Valid Email",
			userEmail: "alice@example.com",
			wantUser: &User{
				ID:     1,
				Name:   "Alice Jones",
				Email:  "alice@example.com",
				Active: true,
			},
			wantError: nil,
		},
		{
			name:      "Non-existent Email",
			userEmail: "doesntexist@example.com",
			wantUser:  nil,
			wantError: ErrRecordNotFound,
		},
	}

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, teardown := newTestDB(t)
			defer teardown()

			m := UserModel{db, infoLog, errorLog}

			user, err := m.GetByEmail(tt.userEmail)

			if err != tt.wantError {
				t.Errorf("want %v; got %s", tt.wantError, err)
			}

			if user != nil {
				if tt.wantUser.ID != user.ID {
					t.Errorf("want %v; got %v", tt.wantUser.ID, user.ID)
				}

				if tt.wantUser.Email != user.Email {
					t.Errorf("want %v; got %v", tt.wantUser.Email, user.Email)
				}
			}
		})
	}
}

func TestUserModelGetForToken(t *testing.T) {
	if testing.Short() {
		t.Skip("mysql: skipping integration test")
	}

	tests := []struct {
		name           string
		tokenPlaintext string
		user           *User
		wantError      error
	}{
		{
			name:           "Valid Email",
			tokenPlaintext: "YER2CR34UPN5IICUGMZAFRA2VA",
			user: &User{
				ID:     1,
				Name:   "Alice Jones",
				Email:  "alice@example.com",
				Active: true,
			},
			wantError: nil,
		},
		{
			name:           "Non-existent Email",
			tokenPlaintext: "face_token",
			user:           nil,
			wantError:      ErrRecordNotFound,
		},
	}

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, teardown := newTestDB(t)
			defer teardown()

			m := UserModel{db, infoLog, errorLog}

			user, err := m.GetForToken(ScopeAuthentication, tt.tokenPlaintext)

			if err != tt.wantError {
				t.Errorf("want %v; got %s", tt.wantError, err)
			}

			if user != nil {
				if tt.user.ID != user.ID {
					t.Errorf("want %v; got %v", tt.user.ID, user.ID)
				}

				if tt.user.Email != user.Email {
					t.Errorf("want %v; got %v", tt.user.Email, user.Email)
				}
			}
		})
	}
}
