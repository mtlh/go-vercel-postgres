package api

import (
	"encoding/json"
	"go_vercel_test/types"
	"net/http"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {

	msg := types.Message{
		Message: "Hello!",
		Routes: map[string]string{
			"/api":              "You are here.",
			"/api/date":         "Current date.",
			"/api/dbget":        "Return some all USERS table information from a connected postgres db.",
			"/api/dbinsert/x":   "Insert into USERS table with name (x).",
			"/api/dbdelete/y":   "Delete from USERS table with id (y).",
			"/api/dbupdate/y/x": "Update from USERS table with id (y) and name (x).",
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
