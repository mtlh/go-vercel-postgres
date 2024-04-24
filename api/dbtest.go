package api

import (
	"encoding/json"
	"go_vercel_test/db"
	"net/http"

	_ "github.com/lib/pq"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func DBTestHandler(w http.ResponseWriter, r *http.Request) {

	db.InitDB()
	defer db.GetDB().Close()
	rows, err := db.GetDB().Query("SELECT id, name FROM users")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()
	// Process query results
	var users []User
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.ID, &user.Name); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		users = append(users, user)
	}
	// Check for errors during row iteration
	if err := rows.Err(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// Marshal users slice to JSON
	jsonData, err := json.Marshal(users)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// Set Content-Type header
	w.Header().Set("Content-Type", "application/json")
	// Write JSON response
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}
