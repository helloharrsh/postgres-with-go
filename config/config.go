package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// LoadEnv loads environment variables once at startup
func LoadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("No .env file found, using system env")
	}
}

// GetEnv safely fetches env values
func GetEnv(key string) string {
	return os.Getenv(key)
}
