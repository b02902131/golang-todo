package main

import (
	"database/sql"
	"golang-todo/handlers"

	"github.com/gorilla/mux"
)

// NewRouter creates and configures a new router
func NewRouter(db *sql.DB) *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/todos", handlers.GetTodosHandler(db)).Methods("GET")
	r.HandleFunc("/todos", handlers.CreateTodoHandler(db)).Methods("POST")
	r.HandleFunc("/todos/{id}", handlers.UpdateTodoHandler(db)).Methods("PUT")
	r.HandleFunc("/todos/{id}", handlers.DeleteTodoHandler(db)).Methods("DELETE")
	// 添加更多路由配置
	return r
}