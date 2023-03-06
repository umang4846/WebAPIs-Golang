package main

import (
	"fmt"
	"net/http"
)

func (app *application) logError(r *http.Request, err error) {
	app.logger.Println(err)
}

//We want to send json formatted error message
func (app *application) errorResponse(w http.ResponseWriter, r *http.Request, status int, message interface{}) {
	//Create the JSON response
	env := envelope{"error": message}
	err := app.writeJSON(w, status, env, nil)
	if err != nil {
		app.logError(r, err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}

//server error response
func (app *application) serverErrorResponse(w http.ResponseWriter, r *http.Request, err error) {
	//We log the error
	app.logError(r, err)
	//Prepare a message with the error
	message := "the server encountered a problem  and could not process the request"
	app.errorResponse(w, r, http.StatusInternalServerError, message)
}

//The not found response
func (app *application) notFoundResponse(w http.ResponseWriter, r *http.Request) {
	//Prepare a message
	message := "the requested resource could not be found"
	app.errorResponse(w, r, http.StatusNotFound, message)
}

//A method not allowed response
func (app *application) methodNotAllowedResponse(w http.ResponseWriter, r *http.Request) {
	//Prepare a message
	message := fmt.Sprintf("the %s method is not supported for this resource", r.Method)
	app.errorResponse(w, r, http.StatusMethodNotAllowed, message)

}
