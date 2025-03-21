package router

import (
	"net/http"
	"github.com/amreshpro/go-todo/internal/controller"
)

var Router http.Server

func TodoRouter() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/",controller.HealthCheck)

	mux.HandleFunc("/todos/", func(w http.ResponseWriter, r *http.Request) {
        switch r.Method {
        case http.MethodGet:
            controller.GetTodoByID(w, r)
        case http.MethodDelete:
            controller.DeleteTodo(w, r)
        default:
            http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        }
    })

 mux.HandleFunc("/todos", func(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		controller.GetTodos(w, r)
	} else if r.Method == http.MethodPost {
		controller.CreateTodo(w, r)
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
})


	return mux
}
