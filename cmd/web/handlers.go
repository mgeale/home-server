package main

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/mgeale/homeserver/pkg/models"
)

func (app *application) getLatestBalances(w http.ResponseWriter, r *http.Request) {
	b, err := app.balances.Latest()
	if err != nil {
		app.serverError(w, err)
		return
	}

	json.NewEncoder(w).Encode(b)
}

func (app *application) getBalance(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get(":id"))
	if err != nil || id < 1 {
		app.notFound(w)
		return
	}

	b, err := app.balances.Get(id)
	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			app.notFound(w)
		} else {
			app.serverError(w, err)
		}
		return
	}

	json.NewEncoder(w).Encode(b)
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

func (app *application) getTransactions(w http.ResponseWriter, r *http.Request) {
	ts, err := app.transaction.Latest()
	if err != nil {
		app.serverError(w, err)
		return
	}

	json.NewEncoder(w).Encode(ts)
}

func (app *application) getTransaction(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get(":id"))
	if err != nil || id < 1 {
		app.notFound(w)
		return
	}

	t, err := app.transaction.Get(id)
	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			app.notFound(w)
		} else {
			app.serverError(w, err)
		}
		return
	}

	json.NewEncoder(w).Encode(t)
}

func (app *application) createTransaction(w http.ResponseWriter, r *http.Request) {
	var t *models.Transaction

	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	_, err = app.transaction.Insert(t.Name, t.Amount, t.Date, t.Type)
	if err != nil {
		app.serverError(w, err)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func ping(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("OK"))
}
