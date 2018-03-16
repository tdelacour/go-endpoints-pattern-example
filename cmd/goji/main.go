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

func helloMiddleware(inner http.Handler) http.Handler {
	mw := func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Do a hello thing\n")
		inner.ServeHTTP(w, r)
		log.Printf("Finish hello thing and inner\n")
	}
	return http.HandlerFunc(mw)
}

func goodbyeMiddleware(inner http.Handler) http.Handler {
	mw := func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Do a goodbye thing\n")
		inner.ServeHTTP(w, r)
		log.Printf("Finish goodbye thing and inner\n")
	}
	return http.HandlerFunc(mw)
}

func hello(w http.ResponseWriter, r *http.Request) {
	name := pat.Param(r, "name")
	log.Printf("\n\n\n hello, %s\n\n\n", name)
}

func helloPost(w http.ResponseWriter, r *http.Request) {
	name := pat.Param(r, "name")
	log.Printf("\n\n\n hello, %s poster\n\n\n", name)
}

func goodbye(w http.ResponseWriter, r *http.Request) {
	first := pat.Param(r, "first")
	last := pat.Param(r, "last")
	log.Printf("\n\n\n goodbye, %s %s\n\n\n", first, last)
}

func goodbyePost(w http.ResponseWriter, r *http.Request) {
	first := pat.Param(r, "first")
	last := pat.Param(r, "last")
	log.Printf("\n\n\n goodbye, %s %s poster\n\n\n", first, last)
}

func main() {
	mux := goji.NewMux()
	mux.Use(fooMiddleware)
	mux.Use(barMiddleware)

	helloMux := goji.SubMux()
	helloMux.Use(helloMiddleware)

	goodbyeMux := goji.SubMux()
	goodbyeMux.Use(goodbyeMiddleware)

	mux.Handle(pat.New("/hello/*"), helloMux)
	mux.Handle(pat.New("/goodbye/*"), goodbyeMux)

	helloMux.HandleFunc(pat.Get("/:name"), hello)
	helloMux.HandleFunc(pat.Post("/:name"), helloPost)

	goodbyeMux.HandleFunc(pat.Get("/:first/:last"), goodbye)
	goodbyeMux.HandleFunc(pat.Post("/:first/:last"), goodbyePost)

	http.ListenAndServe("localhost:8081", mux)
}
