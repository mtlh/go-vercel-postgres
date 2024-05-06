package api

import (
	"go_vercel_test/db"
	"net/http"
	"regexp"
	"strings"
)

func onlyNumbers(s string) string {
	reg := regexp.MustCompile("[^0-9]+")
	cleaned := reg.ReplaceAllString(s, "")
	return cleaned
}

func DBDeleteHandler(w http.ResponseWriter, r *http.Request) {

	parts := strings.Split(r.URL.Path, "/")
	var variable string

	if len(parts) != 4 || parts[3] == "" {
		http.Error(w, "ID not found in URL -> follow /api/dbdelete/id", http.StatusBadRequest)
		return
	} else {
		variable = onlyNumbers(parts[3])
	}

	db.InitDB()
	defer db.GetDB().Close()
	_, err := db.GetDB().Exec("DELETE FROM USERS WHERE id=$1", (variable))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(variable + " removed from db :)"))
}
