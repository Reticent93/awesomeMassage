package main

import (
	"awesomeMassage/internal/repository"
	"awesomeMassage/internal/repository/dbrepo"
	"flag"
	"log"
	"net/http"
)

const portNumber = ":8080"

type application struct {
	Domain string
	DSN    string
	DB     repository.DatabaseRepo
}

var app application

func main() {

	flag.StringVar(&app.DSN, "dsn", "host=localhost port=5433 user=postgres password=postgres dbname=massage sslmode=disable timezone=UTC connect_timeout=5", "Postgres connection string")

	flag.Parse()

	//connect to the database
	conn, err := app.connectToDB()
	if err != nil {
		log.Fatal(err)
	}

	app.DB = &dbrepo.PostgresDBRepo{DB: conn}

	log.Println("Starting application on", portNumber)
	err = http.ListenAndServe(portNumber, routes(&app))
	if err != nil {
		panic(err)
	}
}
