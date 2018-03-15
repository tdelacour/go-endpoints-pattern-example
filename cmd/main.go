package main

import (
	"github.com/tdelacour/go-endpoints-pattern-example/internal/animals"
	"github.com/tdelacour/go-endpoints-pattern-example/internal/api"
	"github.com/tdelacour/go-endpoints-pattern-example/internal/cities"
	"github.com/tdelacour/go-endpoints-pattern-example/internal/plants"

	"log"
	"net/http"
	"time"
)

// New APIs get added here
var apis = []api.Api{
	animals.AnimalsT{},
	plants.PlantsT{},
	cities.CitiesT{},
}

func registerEndpoints() {
	for _, api := range apis {
		api.RegisterEndpoints()
	}
}

func main() {
	registerEndpoints()

	s := &http.Server{
		Addr:              ":8081",
		Handler:           nil,              // defaults to http.DefaultServerMux
		ReadTimeout:       10 * time.Second, // for example
		ReadHeaderTimeout: 10 * time.Second,
		WriteTimeout:      10 * time.Second,
		TLSConfig:         nil, // Relevant for SSL
		MaxHeaderBytes:    0,   // defaults to http.DefaultMaxHeaderBytes (1mb)
	}

	log.Fatal(s.ListenAndServe())
}
