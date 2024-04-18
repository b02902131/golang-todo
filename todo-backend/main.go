package main

import (
	"log"
	"net/http"

	"golang-todo/config"
	"golang-todo/db"
)

func main() {
	db := db.InitDB()
	defer db.Close()

	router := NewRouter(db)

	log.Println("Server is running on http://localhost:8080")
	http.Handle("/", config.SetupCORS()(router))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
