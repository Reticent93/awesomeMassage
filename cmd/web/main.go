package main

import (
	"log"
	"net/http"
)

const portNumber = ":8080"

type application struct {
	Domain string
	DSN    string
}

var app application

func main() {
	
	log.Println("Starting application on", portNumber)
	err := http.ListenAndServe(portNumber, routes(&app))
	if err != nil {
		panic(err)
	}
}
