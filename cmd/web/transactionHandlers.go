package main

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/mgeale/homeserver/pkg/models"
)

func (app *application) getTransactions(w http.ResponseWriter, r *http.Request) {
	ts, err := app.transactions.Latest()
	if err != nil {
		app.serverError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(ts)
}

func (app *application) getTransaction(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get(":id"))
	if err != nil || id < 1 {
		app.notFound(w)
		return
	}

	t, err := app.transactions.Get(id)
	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			app.notFound(w)
		} else {
			app.serverError(w, err)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(t)
}

func (app *application) createTransaction(w http.ResponseWriter, r *http.Request) {
	var t *models.Transaction

	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	_, err = app.transactions.Insert(t.Name, t.Amount, t.Date, t.Type)
	if err != nil {
		app.serverError(w, err)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (app *application) updateTransaction(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get(":id"))
	if err != nil || id < 1 {
		app.notFound(w)
		return
	}

	var t *models.Transaction

	err = json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	if t == nil || t.Name == "" || t.Amount == 0 || t.Date == "" || t.Type == "" {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	err = app.transactions.Update(id, t.Name, t.Amount, t.Date, t.Type)
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

func (app *application) deleteTransaction(w http.ResponseWriter, r *http.Request) {
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
