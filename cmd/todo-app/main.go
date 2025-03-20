package main

import (
	"net/http"

	"github.com/amreshpro/go-todo/config"
	"github.com/amreshpro/go-todo/internal/router"
	"github.com/amreshpro/go-todo/pkg/logger"
)

func main() {
	// initialize logger
	logger.InitLogger()
	config.MustLoad()

	// initialize router
	todoRouter := router.TodoRouter()

	// start server
	logger.Log.Info("🚀 Server started at http://localhost:" + config.PORT)
	err := http.ListenAndServe(":"+config.PORT, todoRouter)
	if err != nil {
		logger.Log.Error("❌ Failed to start server")
	}
}
