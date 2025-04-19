/*
Routes handles the handlers
	Notfound - When a requested handler is ot available
	MethonotAllowed - When the method is not allowed on the handler
Handlers handles the requests to the server
	Notfound - When the requested response is not found
*/

package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() http.Handler {
	router := httprouter.New()

	// Convert the notFoundResponse() helper to a http.Handler using the
	// http.HandlerFunc() adapter, and then set it as the custom error handler for 404
	// Not Found responses.
	router.NotFound = http.HandlerFunc(app.notFoundResponse)
	// Likewise, convert the methodNotAllowedResponse() helper to a http.Handler and set
	// it as the custom error handler for 405 Method Not Allowed responses.
	router.MethodNotAllowed = http.HandlerFunc(app.methodNotAllowedResponse)

	router.HandlerFunc(http.MethodGet, "/v1/healthcheck", app.healthcheckHandler)
	router.HandlerFunc(http.MethodPost, "/v1/games", app.createGameHandler)
	router.HandlerFunc(http.MethodGet, "/v1/games/:id", app.showGameHandler)

	return app.recoverPanic(router)
}
