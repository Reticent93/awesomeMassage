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
func (app *application) jsonResponse(w http.ResponseWriter, r *http.Request, status int, data interface{}, headers ...http.Header) error {
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
