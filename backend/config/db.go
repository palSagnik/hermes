package config

import (
	"os"

	_ "github.com/joho/godotenv"
)

var DB_USER = os.Getenv("POSTGRES_USER")
var DB_HOST = os.Getenv("POSTGRES_HOST")
var DB_PASS = os.Getenv("POSTGRES_PASSWORD")
var DB_NAME = os.Getenv("POSTGRES_DATABASE")
var SSL_MODE = os.Getenv("POSTGRES_SSLMODE") 
var DB_PORT = 5432
