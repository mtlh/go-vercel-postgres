package api

import (
	"go_vercel_test/db"
	"net/http"
	"strings"
)

func DBUpdateHandler(w http.ResponseWriter, r *http.Request) {

	parts := strings.Split(r.URL.Path, "/")
	var namevar string
	var idvar string

	if len(parts) != 5 || parts[3] == "" || parts[4] == "" {
		http.Error(w, "ID and Name not found in URL -> follow /api/dbupdate/id/name", http.StatusBadRequest)
		return
	} else {
		idvar = onlyNumbers(parts[3])
		namevar = cleanString(parts[4])
	}

	db.InitDB()
	defer db.GetDB().Close()
	_, err := db.GetDB().Exec("UPDATE USERS SET name=$2 WHERE id=$1", idvar, namevar)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(namevar + " updated for id " + idvar + " in db :)"))
}
