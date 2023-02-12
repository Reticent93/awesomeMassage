package main

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
	"strconv"
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

	_ = app.jsonResponse(w, http.StatusOK, payload)
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

func (app *application) GetTherapists(w http.ResponseWriter, r *http.Request) {
	therapists, err := app.DB.GetTherapists()
	if err != nil {
		app.errorJSON(w, err)
		return
	}
	_ = app.jsonResponse(w, http.StatusOK, therapists)
}

func (app *application) GetATherapist(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	therapistID, err := strconv.Atoi(id)
	if err != nil {
		app.errorJSON(w, err)
		return
	}
	therapist, err := app.DB.OneTherapist(therapistID)
	if err != nil {
		app.errorJSON(w, err)
		return
	}
	_ = app.jsonResponse(w, http.StatusOK, therapist)
}

func (app *application) AllUsers(w http.ResponseWriter, r *http.Request) {
	users, err := app.DB.GetUsers()
	if err != nil {
		app.errorJSON(w, err)
		return
	}
	_ = app.jsonResponse(w, http.StatusOK, users)
}

func (app *application) GetAUser(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	userID, err := strconv.Atoi(id)
	if err != nil {
		app.errorJSON(w, err)
		return
	}
	user, err := app.DB.OneUser(userID)

	if err != nil {
		app.errorJSON(w, err)
		return
	}
	_ = app.jsonResponse(w, http.StatusOK, user)
}
