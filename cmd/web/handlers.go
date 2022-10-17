package main

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/mgeale/homeserver/pkg/models"
)

func (app *application) getLatestBalances(w http.ResponseWriter, r *http.Request) {
	s, err := app.balances.Latest()
	if err != nil {
		app.serverError(w, err)
		return
	}

	json.NewEncoder(w).Encode(s)
}

func (app *application) getBalance(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get(":id"))
	if err != nil || id < 1 {
		app.notFound(w)
		return
	}

	s, err := app.balances.Get(id)
	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			app.notFound(w)
		} else {
			app.serverError(w, err)
		}
		return
	}

	json.NewEncoder(w).Encode(s)
}

func (app *application) createBalance(w http.ResponseWriter, r *http.Request) {
	var b *models.Balance

	err := json.NewDecoder(r.Body).Decode(&b)
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	_, err = app.balances.Insert(b.Name, b.Balance, b.BalanceAUD, b.PricebookID, b.ProductID)
	if err != nil {
		app.serverError(w, err)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func ping(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("OK"))
}
