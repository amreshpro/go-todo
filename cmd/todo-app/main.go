package main

import (
	
	"net/http"

	"github.com/amreshpro/go-todo/config"
	"github.com/amreshpro/go-todo/internal/router"
	"github.com/amreshpro/go-todo/pkg/logger"
	"github.com/joho/godotenv"
)

func main(){
	
	// load env
	if err:= godotenv.Load();err!=nil{
		logger.Log.Error("Failed to load environment variable or .env not found")
	}
	// load logger
	logger.InitLogger()

	// load load db

	// router
	todoRouter:=router.TodoRouter()

	// server
	logger.Log.Info("Server started at http://localhost:"+config.PORT)
	err:=http.ListenAndServe(":"+config.PORT,todoRouter)
	if err!=nil {
		logger.Log.Error("Failed to started server")
	}
	

}
