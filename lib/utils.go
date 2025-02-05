package utils

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading env files")
	}
}

func GetApiKey() string {
	return os.Getenv("GEMINI_API_KEY")
}
