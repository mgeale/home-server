package mock

import (
	"time"

	"github.com/mgeale/homeserver/internal/db"
)

var mockToken = &db.Token{
	Plaintext: "plaintext",
	Hash:      []byte("hash"),
	UserID:    1,
	Expiry:    time.Now(),
	Scope:     "admin",
}

type TokenModel struct{}

func (t *TokenModel) DeleteAllForUser(scope string, userID int64) error {
	return nil
}
func (t *TokenModel) Insert(token *db.Token) error {
	return nil
}
func (t *TokenModel) New(userID int64, ttl time.Duration, scope string) (*db.Token, error) {
	return mockToken, nil
}
