package app

import (
	"context"
	"net/http"

	"github.com/mgeale/homeserver/internal/db"
)

type contextKey string

const userContextKey = contextKey("user")

func (app *Application) contextSetUser(r *http.Request, user *db.User) *http.Request {
	ctx := context.WithValue(r.Context(), userContextKey, user)
	return r.WithContext(ctx)
}

func (app *Application) contextGetUser(r *http.Request) *db.User {
	user, ok := r.Context().Value(userContextKey).(*db.User)
	if !ok {
		panic("missing user value in request context")
	}

	return user
}
