package config

import (
	"fmt"
	"os"

	"github.com/amreshpro/go-todo/pkg/logger"
	"github.com/joho/godotenv"
)

var (
	PORT string
)

func MustLoad() {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		logger.InitLogger()
		logger.Log.Error("❌ Failed to load .env file or .env not found")
	}

	// Now safely read from env
	PORT = getEnv("PORT", "5555")
}

func getEnv(key string, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		fmt.Println("✅ Loaded from .env:", key, "=", value)
		return value
	}
	fmt.Println("⚠️ Using fallback for:", key)
	return fallback
}
