package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/xKachi/gamify.git/internal/data"
)

func (app *application) createGameHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "create a new game")
}

func (app *application) showGameHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)

	if err != nil || id < 1 {
		app.notFoundResponse(w, r)
		return
	}

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
