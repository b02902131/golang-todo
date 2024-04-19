package main

import (
	"database/sql"
	"golang-todo/handlers"
	"log"

	"github.com/gorilla/mux"
)

// NewRouter creates and configures a new router
func NewRouter(db *sql.DB) *mux.Router {
	log.Println("Creating new router")
	r := mux.NewRouter()
	r.HandleFunc("/", handlers.HelloWorld).Methods("GET")
	r.HandleFunc("/todos", handlers.GetTodosHandler(db)).Methods("GET")
	r.HandleFunc("/todos", handlers.CreateTodoHandler(db)).Methods("POST")
	r.HandleFunc("/todos/{id}", handlers.UpdateTodoHandler(db)).Methods("PUT")
	r.HandleFunc("/todos/{id}", handlers.DeleteTodoHandler(db)).Methods("DELETE")
	// 添加更多路由配置
	log.Println("Router created")
	return r
}
