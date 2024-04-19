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

	log.Println("Setting up server CORS...")
	http.Handle("/", config.SetupCORS()(router))

	log.Println("Starting server...")
	log.Fatal(http.ListenAndServe("0.0.0.0:8080", nil))
	log.Println("Listening on :8080")
}
