package animals

import (
	"encoding/json"
	"log"
	"net/http"
)

type AnimalsT struct{} // Empty struct to enforce type

/*************************************
 * PIG REQUESTS
 *************************************/

func pig(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		log.Printf("POST to /animals/pig")
		w.WriteHeader(http.StatusOK)
	case http.MethodGet:
		log.Printf("GET to /animals/pig")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]interface{}{"sound": "oink oink"})
	case http.MethodPut:
		log.Printf("PUT to /animals/pig")
		w.WriteHeader(http.StatusOK)
	// New request types get added here
	default:
		http.NotFound(w, r) // 404
	}
}

/*************************************
 * COW REQUESTS
 *************************************/

func cow(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		log.Printf("GET to /animals/cow")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]interface{}{"sound": "moooooo"})
		return
	}

	http.NotFound(w, r)
}

/*************************************
 * HORSE REQUESTS
 *************************************/

func horse(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		log.Printf("GET to /animals/horse")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]interface{}{"sound": "neighhhh"})
		return
	}

	http.NotFound(w, r)
}

// This is the only public function!
// New endpoints get added here
func (a AnimalsT) RegisterEndpoints() { // if it walks like a duck...
	http.HandleFunc("/animals/pig", pig)
	http.HandleFunc("/animals/cow", cow)
	http.HandleFunc("/animals/horse", horse)
}
