package controller

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/amreshpro/go-todo/internal/model"
)

var todos = []model.Todo{}
var idCounter int64 = 1

// CreateTodo handles POST /todos
func CreateTodo(w http.ResponseWriter, r *http.Request) {
    var todo model.Todo

    err := json.NewDecoder(r.Body).Decode(&todo)
    if err != nil {
        http.Error(w, "Invalid request body", http.StatusBadRequest)
        return
    }

    todo.ID = idCounter
    idCounter++
    todo.CreatedAt = time.Now().UTC()
    todo.UpdatedAt = time.Now().UTC()
    todos = append(todos, todo)

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(todo)
}

// GetTodos handles GET /todos
func GetTodos(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(todos)
}

// GetTodoByID handles GET /todos/{id}
func GetTodoByID(w http.ResponseWriter, r *http.Request) {
    idStr := strings.TrimPrefix(r.URL.Path, "/todos/")
    id, err := strconv.ParseInt(idStr, 10, 64)
    if err != nil {
        http.Error(w, "Invalid ID", http.StatusBadRequest)
        return
    }

    for _, todo := range todos {
        if todo.ID == id {
            w.Header().Set("Content-Type", "application/json")
            json.NewEncoder(w).Encode(todo)
            return
        }
    }

    http.Error(w, "Todo not found", http.StatusNotFound)
}

// DeleteTodo handles DELETE /todos/{id}
func DeleteTodo(w http.ResponseWriter, r *http.Request) {
    idStr := strings.TrimPrefix(r.URL.Path, "/todos/")
    id, err := strconv.ParseInt(idStr, 10, 64)
    if err != nil {
        http.Error(w, "Invalid ID", http.StatusBadRequest)
        return
    }

    for i, todo := range todos {
        if todo.ID == id {
            todos = append(todos[:i], todos[i+1:]...)
            w.WriteHeader(http.StatusNoContent)
            return
        }
    }

    http.Error(w, "Todo not found", http.StatusNotFound)
}
