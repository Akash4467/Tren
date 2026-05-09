package main

import (
	"database/sql"
	"fmt"
	"net/http"
)

func getUser(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	id := r.URL.Query().Get("id")

	// Vulnerable SQL injection
	query := "SELECT * FROM users WHERE id = " + id

	rows, err := db.Query(query)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return 
	}

	defer rows.Close()
 
	fmt.Fprintf(w, "User fetched")
}

func main() {
	fmt.Println("server started")
}
