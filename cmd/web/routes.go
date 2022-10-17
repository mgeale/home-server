package main

import (
	"net/http"

	"github.com/bmizerany/pat"
)

func (app *application) routes() http.Handler {
	mux := pat.New()
	mux.Post("/balance/create", app.basicAuth(http.HandlerFunc(app.createBalance)))
	mux.Get("/balance/:id", app.basicAuth(http.HandlerFunc(app.getBalance)))
	mux.Get("/balances", app.basicAuth(http.HandlerFunc(app.getLatestBalances)))

	mux.Get("/ping", http.HandlerFunc(ping))

	return app.recoverPanic(app.logRequest(secureHeaders(mux)))
}
