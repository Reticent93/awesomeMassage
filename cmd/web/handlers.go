package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func (app *application) Home(w http.ResponseWriter, r *http.Request) {
	
	var payload = struct {
		Status  string `json:"status"`
		Message string `json:"message"`
		Version string `json:"version"`
	}{
		Status:  "active",
		Message: "Welcome to the Message Board API",
		Version: "1.0.0",
	}
	
	_ = app.jsonResponse(w, r, http.StatusOK, payload)
}

func (app *application) About(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["test"] = "This is the about."
	
	output, err := json.Marshal(stringMap)
	if err != nil {
		log.Println(err)
	}
	w.Header().Set("Content-Type", "application/json")
	//w.WriteHeader(http.StatusOK)
	_, err = w.Write(output)
	
}
