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

	w.Header().Set("Content-Type", "application/json")
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

	w.Header().Set("Content-Type", "application/json")
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

func (app *application) updateBalance(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get(":id"))
	if err != nil || id < 1 {
		app.notFound(w)
		return
	}

	var b *models.Balance

	err = json.NewDecoder(r.Body).Decode(&b)
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	if b == nil || b.Name == "" || b.Balance == 0 || b.BalanceAUD == 0 || b.PricebookID == 0 || b.ProductID == 0 {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	err = app.balances.Update(id, b.Name, b.Balance, b.BalanceAUD, b.PricebookID, b.ProductID)
	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			app.notFound(w)
		} else {
			app.serverError(w, err)
		}
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (app *application) deleteBalance(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get(":id"))
	if err != nil || id < 1 {
		app.notFound(w)
		return
	}

	err = app.transactions.Delete(id)
	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			app.notFound(w)
		} else {
			app.serverError(w, err)
		}
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
