package api

import (
	"go_vercel_test/db"
	"net/http"
	"regexp"
	"strings"
)

func cleanString(s string) string {
	reg := regexp.MustCompile("[^a-zA-Z0-9]+")
	cleaned := reg.ReplaceAllString(s, "")
	return cleaned
}

func DBInsertHandler(w http.ResponseWriter, r *http.Request) {

	parts := strings.Split(r.URL.Path, "/")
	var variable string

	if len(parts) != 4 || parts[3] == "" {
		http.Error(w, "Name not found in URL -> follow /api/dbinsert/name", http.StatusBadRequest)
		return
	} else {
		variable = cleanString(parts[3])
	}

	db.InitDB()
	defer db.GetDB().Close()
	_, err := db.GetDB().Exec("INSERT INTO USERS(name) VALUES($1)", variable)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(variable + " added to db :)"))
}
