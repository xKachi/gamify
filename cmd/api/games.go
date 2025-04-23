package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/xKachi/gamify.git/internal/data"
)

func (app *application) createGameHandler(w http.ResponseWriter, r *http.Request) {
	// struct for decoding the request from the client
	var input struct {
		Title       string   `json:"title"`
		Developer   string   `json:"developer"`
		ReleaseYear int32    `json:"release_year"`
		Genres      []string `json:"genres"`
		Platforms   []string `json:"platforms"`
	}

	// Use the new readJSON() helper to decode the request body into the input struct.
	// If this returns an error we send the client the error message along with a 400
	// Bad Request status code, just like before.
	err := app.readJSON(r, w, &input)

	if err != nil {
		app.errorResponse(w, r, http.StatusBadRequest, err.Error())
		return
	}
	fmt.Fprintf(w, "%+v\n", input)
}

func (app *application) showGameHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)

	if err != nil || id < 1 {
		app.notFoundResponse(w, r)
		return
	}

	// struct we expected to be in the response body, i.e the structure of the response
	game := data.Game{
		ID:          id,
		ReleaseYear: 2017,
		CreatedAt:   time.Now(),
		Title:       "The Legend of Zelda: Breath of the Wild",
		Developer:   "Nintendo",
		Genres:      []string{"action", "adventure"},
		Platforms:   []string{"Nintedo Switch", "PS5"},
		Version:     1,
	}

	err = app.WriteJSON(w, http.StatusOK, envelope{"game": game}, nil)

	if err != nil {
		app.serverResponse(w, r, err)
	}
}
