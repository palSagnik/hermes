package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/gofiber/fiber/v2/log"
	"github.com/palSagnik/daily-expenses-application/config"
)

var (
	DB *sql.DB
	err error
)

// connecting to the database
func ConnectDB() error {
	
	connectionString := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s", config.DB_HOST, config.DB_USER, config.DB_PASS, config.DB_NAME, config.DB_PORT, config.SSL_MODE)
	DB, err = sql.Open("postgres", connectionString)
	if err != nil {
		return err
	}

	err = DB.Ping()
	if err != nil {
		return err
	}

	log.Infof("connected to database: %s\n", config.DB_NAME)
	
	return nil
}
