package main

import (
	"log"
	"net/http"

	"golang-todo/db"
)

func main() {
	db := db.InitDB()
	defer db.Close()

	router := NewRouter(db)

	log.Println("Server is running on http://localhost:8080")
	http.ListenAndServe(":8080", router)

}
