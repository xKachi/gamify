/**
We mainly have two types of errors
1. Application
2. Server/Client Errors

For errors handled mainly on the application we only send application errors.
For errors handled on the server, we can send neither or both server and application errors.
**/

package main

import (
	"fmt"
	"net/http"
)

// application error
// The logError() method is a generic helper for logging an error message along
// with the current request method and URL as attributes in the log entry
func (app *application) logError(r *http.Request, err error) {
	var (
		method = r.Method           // GET, POST
		uri    = r.URL.RequestURI() // Request path
	)

	app.logger.Error(err.Error(), "method", method, "uri", uri)
}

// server/client error
// messages to the client with a given status code. Note that we're using the any
// type for the message parameter, rather than just a string type, as this gives us
// more flexibility over the values that we can include in the response.
func (app *application) errorResponse(w http.ResponseWriter, r *http.Request, status int, message any) {
	env := envelope{"error": message}

	err := app.WriteJSON(w, status, env, nil)

	if err != nil {
		// app error
		app.logError(r, err)
		// srv error
		w.WriteHeader(500)
	}
}

// The serverErrorResponse() method will be used when our application encounters an
// unexpected problem at runtime. It logs the detailed error message, then uses the
// errorResponse() helper to send a 500 Internal Server Error status code and JSON
// response (containing a generic error message) to the client.
func (app *application) serverResponse(w http.ResponseWriter, r *http.Request, err error) {
	app.logError(r, err)

	message := "the server encountered a problem and could not process your request"

	app.errorResponse(w, r, http.StatusInternalServerError, message)
}

// The notFoundResponse() method will be used to send a 404 Not Found status code and
// JSON response to the client.
func (app *application) notFoundResponse(w http.ResponseWriter, r *http.Request) {
	message := "the requested source could not be found"
	app.errorResponse(w, r, http.StatusNotFound, message)
}

// The methodNotAllowedResponse() method will be used to send a 405 Method Not Allowed
// status code and JSON response to the client.
func (app *application) methodNotAllowedResponse(w http.ResponseWriter, r *http.Request) {
	messgae := fmt.Sprintf("the %s method is not supported for this resource", r.Method)
	app.errorResponse(w, r, http.StatusMethodNotAllowed, messgae)
}
