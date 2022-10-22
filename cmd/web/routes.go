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

	mux.Post("/transaction/create", app.basicAuth(http.HandlerFunc(app.createTransaction)))
	mux.Get("/transaction/:id", app.basicAuth(http.HandlerFunc(app.getTransaction)))
	mux.Get("/transactions", app.basicAuth(http.HandlerFunc(app.getTransactions)))

	mux.Get("/ping", http.HandlerFunc(ping))

	return app.recoverPanic(app.logRequest(secureHeaders(mux)))
}
