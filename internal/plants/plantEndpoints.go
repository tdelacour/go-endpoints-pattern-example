package plants

import (
	"encoding/json"
	"log"
	"net/http"
)

type PlantsT struct{} // Empty struct to enforce type

/*************************************
 * TREE REQUESTS
 *************************************/

func tree(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		log.Printf("POST to /plants/tree")
		w.WriteHeader(http.StatusOK)
	case http.MethodGet:
		log.Printf("GET to /plants/tree")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(
			map[string]interface{}{
				"types": []string{"oak", "maple", "pine", "poplar", "redwood"},
			},
		)
	case http.MethodPut:
		log.Printf("PUT to /plants/tree")
		w.WriteHeader(http.StatusOK)
	// New request types get added here
	default:
		http.NotFound(w, r) // 404
	}
}

/*************************************
 * FLOWER REQUESTS
 *************************************/

func flower(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		log.Printf("GET to /plants/flower")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(
			map[string]interface{}{
				"color": "varied",
			},
		)
		return
	}

	http.NotFound(w, r)
}

/*************************************
 * HERB REQUESTS
 *************************************/

func herb(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		log.Printf("GET to /plants/herb")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(
			map[string]interface{}{
				"tastiest": "basil, duh",
			},
		)
		return
	}

	http.NotFound(w, r)
}

// This is the only public function!
// New endpoints get added here
func (p PlantsT) RegisterEndpoints() { // if it walks like a duck...
	http.HandleFunc("/plants/tree", tree)
	http.HandleFunc("/plants/flower", flower)
	http.HandleFunc("/plants/herb", herb)
}
