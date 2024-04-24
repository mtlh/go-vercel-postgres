package api

import (
	"encoding/json"
	"net/http"
)

type Message struct {
	Message string            `json:"message"`
	Routes  map[string]string `json:"routes"`
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {

	msg := Message{
		Message: "Hello!",
		Routes: map[string]string{
			"/api":        "You are here.",
			"/api/dbtest": "Return some test information from a connected postgres db.",
		},
	}

	jsonData, err := json.Marshal(msg)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}
