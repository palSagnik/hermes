package config

import (
	"fmt"
	"os"

	_ "github.com/joho/godotenv/autoload"
)

var SESSION_SECRET = os.Getenv("SESSION_SECRET")
var SESSION_EXPIRY = 72
var MAIL_LEN = 320
var PASS_LEN = 48
var NAME_LEN = 24

var APP_PORT = fmt.Sprintf(":%s", os.Getenv("APP_PORT"))