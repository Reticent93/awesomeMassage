package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
)

func routes(app *application) http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Recoverer)
	r.Use(middleware.Logger)
	r.Use(app.enableCORS)

	r.Get("/", app.Home)
	r.Get("/about", app.About)
	r.Get("/therapists", app.GetTherapists)
	r.Get("/therapists/{id}", app.GetATherapist)
	r.Get("/users", app.AllUsers)
	r.Get("/users/{id}", app.GetAUser)

	return r
}
