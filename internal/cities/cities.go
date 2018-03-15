package cities

import (
	"encoding/json"
	"log"
	"net/http"
)

type CitiesT struct{} // Empty struct to enforce type

/*************************************
 * GET AMERICAN CITIES
 *************************************/

func american(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		log.Printf("GET to /cities/american")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(
			[]string{
				"New York",
				"Philadelphia",
				"Washington D.C.",
				"Chicago",
				"Los Angeles",
				"New Orleans",
			},
		)
		return
	}

	http.NotFound(w, r)
}

/*************************************
 * GET EUROPEAN CITIES
 *************************************/

func european(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		log.Printf("GET to /cities/european")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(
			[]string{
				"Paris",
				"Dublin",
				"Milan",
				"Berlin",
				"Krakow",
				"Prague",
				"Brussels",
			},
		)
		return
	}

	http.NotFound(w, r)
}

// This is the only public function!
// New endpoints get added here
func (c CitiesT) RegisterEndpoints() { // if it walks like a duck...
	http.HandleFunc("/cities/american", american)
	http.HandleFunc("/cities/european", european)
}
