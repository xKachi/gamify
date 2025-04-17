package main

import (
	"net/http"
)

func (app *application) healthcheckHandler(w http.ResponseWriter, r *http.Request) {
	data := map[string]string{
		"status":      "available",
		"environment": app.config.env,
		"version":     version,
	}

	err := app.WriteJSON(w, http.StatusOK, data, nil)

	if err != nil {
		// app level error
		app.logger.Error(err.Error())
		// server level error
		http.Error(w, "The server encountered a problem and could not process your request", http.StatusInternalServerError)
		return
	}

}
