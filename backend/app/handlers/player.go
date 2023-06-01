package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/rembrandtcosta/baseball-database/backend/app/database"
)

func PlayerHandler(w http.ResponseWriter, r *http.Request) {
	db, err := database.Connect()
	if err != nil {
		w.WriteHeader(500)
		return
	}
	defer db.Close()

	rows, err := db.Query("SELECT nameFirst FROM players LIMIT 10");
	if err != nil {
		w.WriteHeader(500)
		return
	}
	var names []string
	for rows.Next() {
		var name string
		err = rows.Scan(&name)
		names = append(names, name)
	}
	json.NewEncoder(w).Encode(names)
}


func BlogHandler(w http.ResponseWriter, r *http.Request) {
	db, err := database.Connect()
	if err != nil {
		w.WriteHeader(500)
		return
	}
	defer db.Close()

	rows, err := db.Query("SELECT title FROM blog")
	if err != nil {
		w.WriteHeader(500)
		return
	}
	var titles []string
	for rows.Next() {
		var title string
		err = rows.Scan(&title)
		titles = append(titles, title)
	}
	json.NewEncoder(w).Encode(titles)
}
