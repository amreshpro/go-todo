package config

import "os"


var (
	PORT = getEnv("PORT","8080")
)



func getEnv(key string, fallback string) string {
	if value,ok := os.LookupEnv(key); ok{
		return value
	}
	return fallback
}