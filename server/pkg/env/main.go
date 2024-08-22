package env

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

const (
	DB_DRIVER = iota
	DB_HOST
	DB_PORT
	DB_USER
	DB_PASSWORD
	DB_NAME
	DB_SSLMODE
)

var keys = []string {
	"DB_DRIVER",
	"DB_HOST",
	"DB_PORT",
	"DB_USER",
	"DB_PASSWORD",
	"DB_NAME",
	"DB_SSLMODE",
}

func Load() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file")
	}
}

func Read(key int) string {
	if key < 0 || key >= len(keys) {
		log.Fatalf("Invalid env key")
	}
	return os.Getenv(keys[key])
}
