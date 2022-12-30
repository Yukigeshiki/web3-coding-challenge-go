package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

var (
	Port string
)

// init loads a .env file if the environment is local, then assigns environment variables to config values
func init() {

	if os.Getenv("ENV") == "" {
		if err := godotenv.Load(); err != nil {
			log.Fatal("Error loading .env file")
		}
	}

	Port = os.Getenv("PORT")
}
