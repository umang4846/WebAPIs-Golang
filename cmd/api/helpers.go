package main

import (
	"encoding/json"
	"errors"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
)

type envelope map[string]interface{}

func (app *application) readIDParam(r *http.Request) (int64, error) {
	//Use the "ParamsFromContext()" function to get the request context as a slice
	params := httprouter.ParamsFromContext(r.Context())
	// GET value of the 'id' parameter
	id, err := strconv.ParseInt(params.ByName("id"), 10, 64)
	if err != nil || id < 1 {
		return 0, errors.New("invalid id parameter")
	}
	return id, nil
}
func (app *application) writeJSON(w http.ResponseWriter, status int, data envelope, headers http.Header) error {
	//Convert map into JOSN object
	js, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		return err
	}
	//adding new line in json
	js = append(js, '\n')
	//add the HEADERS
	for key, value := range headers {
		w.Header()[key] = value
	}
	//Specify that we will service our response using JSON
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	//Write the []byte slice containing the JSON response body
	w.Write(js)
	return nil
}
