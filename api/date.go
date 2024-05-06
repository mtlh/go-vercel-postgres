package api

import (
	"net/http"
	"time"
)

func DateHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(time.Now().String()))
}
