package main

import (
	"log"
	"net/http"

	"goji.io"
	"goji.io/pat"
)

// We play with Goji!

func fooMiddleware(inner http.Handler) http.Handler {
	mw := func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Do a foo thing\n")
		inner.ServeHTTP(w, r)
		log.Printf("Finish foo and inner\n")
	}
	return http.HandlerFunc(mw)
}

func barMiddleware(inner http.Handler) http.Handler {
	mw := func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Do a bar thing\n")
		inner.ServeHTTP(w, r)
		log.Printf("Finish bar and inner\n")
	}
	return http.HandlerFunc(mw)
}

func hello(w http.ResponseWriter, r *http.Request) {
	name := pat.Param(r, "name")
	log.Printf("\n\n\n hello, %s\n\n\n", name)
}

func main() {
	mux := goji.NewMux()
	mux.Use(fooMiddleware)
	mux.Use(barMiddleware)
	mux.HandleFunc(pat.Get("/hello/:name"), hello)

	http.ListenAndServe("localhost:8081", mux)
}
