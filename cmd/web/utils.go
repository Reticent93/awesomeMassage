package main

import (
	"encoding/json"
	"net/http"
)

// JSONResponse is a struct that is used to return a JSON response
type JSONResponse struct {
	Error bool   `json:"error"`
	Msg   string `json:"message"`
	Data  string `json:"data,omitempty"`
}

// jsonResponse is a helper function that is used to return a JSON response
func (app *application) jsonResponse(w http.ResponseWriter, status int, data interface{}, headers ...http.Header) error {
	output, err := json.Marshal(data)
	if err != nil {
		return err
	}
	if len(headers) > 0 {
		for key, value := range headers[0] {
			w.Header()[key] = value
		}
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_, err = w.Write(output)
	if err != nil {
		return err
	}
	return nil
}

func (app *application) errorJSON(w http.ResponseWriter, err error, status ...int) error {
	statusCode := http.StatusBadRequest
	if len(status) > 0 {
		statusCode = status[0]
	}
	var payload JSONResponse
	payload.Error = true
	payload.Msg = err.Error()
	return app.jsonResponse(w, statusCode, payload)
}
