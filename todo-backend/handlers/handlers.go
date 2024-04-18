package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"golang-todo/todo"

	"github.com/gorilla/mux"
)

func CreateTodoHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var newTodo todo.Todo
		if err := json.NewDecoder(r.Body).Decode(&newTodo); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// 使用 RETURNING 子句获取新插入行的 ID
		sqlStatement := `INSERT INTO todos (title, description, completed) VALUES ($1, $2, $3) RETURNING id`
		err := db.QueryRow(sqlStatement, newTodo.Title, newTodo.Description, newTodo.Completed).Scan(&newTodo.ID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(newTodo)
	}
}

func GetTodosHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Only GET method is allowed", http.StatusMethodNotAllowed)
			return
		}

		rows, err := db.Query("SELECT id, title, description, completed FROM todos")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		var todos []todo.Todo = []todo.Todo{}
		for rows.Next() {
			var t todo.Todo
			if err := rows.Scan(&t.ID, &t.Title, &t.Description, &t.Completed); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			todos = append(todos, t)
		}

		w.Header().Set("Content-Type", "application/json ")
		json.NewEncoder(w).Encode(todos)
	}
}

func UpdateTodoHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPut {
			http.Error(w, "Only PUT method is allowed", http.StatusMethodNotAllowed)
			return
		}

		vars := mux.Vars(r)
		id := vars["id"]
		var todo todo.Todo
		if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		_, err := db.Exec("UPDATE todos SET title = $1, description = $2, completed = $3 WHERE id = $4", todo.Title, todo.Description, todo.Completed, id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(todo)
	}
}

func DeleteTodoHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodDelete {
			http.Error(w, "Only DELETE method is allowed", http.StatusMethodNotAllowed)
			return
		}

		vars := mux.Vars(r)
		id := vars["id"]

		_, err := db.Exec("DELETE FROM todos WHERE id = $1", id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}
