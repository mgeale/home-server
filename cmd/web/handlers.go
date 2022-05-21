package main

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/mgeale/homeserver/pkg/models"
)

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

	outputBytes, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		log.Printf("Failed to format data")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(outputBytes)
}

func (app *application) createBalance(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	// _, err := app.balances.Insert("title", "content", "expires")
	if err != nil {
		app.serverError(w, err)
		return
	}

	w.Write([]byte("OK"))
}

func ping(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("OK"))
}
