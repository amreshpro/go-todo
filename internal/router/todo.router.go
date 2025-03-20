package router

import (
	"net/http"

	"github.com/amreshpro/go-todo/internal/controller"
)

var Router http.Server

func TodoRouter() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/",controller.HealthCheck)
	return mux
}
