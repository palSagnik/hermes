package database

import (
	"database/sql"

	"github.com/gofiber/fiber/v2/log"
	_ "github.com/lib/pq"
)

func createTables(db *sql.DB) error {
	// Create users table
	userTable := `
	CREATE TABLE IF NOT EXISTS users (
		user_id SERIAL PRIMARY KEY,
		email VARCHAR(255) UNIQUE NOT NULL,
		name VARCHAR(255) NOT NULL,
		number VARCHAR(20) NOT NULL,
		password VARCHAR(255) NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);`

	_, err := db.Exec(userTable)
	if err != nil {
		log.Warn(err)
		return err
	}

	// Create verification table
	verificationTable := `
	CREATE TABLE IF NOT EXISTS verifications (
		verification_id SERIAL PRIMARY KEY,
		email VARCHAR(255) UNIQUE NOT NULL,
		name VARCHAR(255) NOT NULL,
		number VARCHAR(20) NOT NULL,
		password VARCHAR(255) NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);`

	_, err = db.Exec(verificationTable)
	if err != nil {
		log.Warn(err)
		return err
	}

	// Create expenses table
	expenseTable := `
	CREATE TABLE IF NOT EXISTS expenses (
		expense_id SERIAL PRIMARY KEY,
		user_id INTEGER NOT NULL,
		amount DECIMAL(10,2) NOT NULL,
		description TEXT NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		FOREIGN KEY (user_id) REFERENCES users(user_id) ON DELETE CASCADE
	);`

	_, err = db.Exec(expenseTable)
	if err != nil {
		log.Warn(err)
		return err
	}

	log.Info("all tables created successfully")
	return nil
}

func MigrateUp() error {
	
	err := createTables(DB)
	if err != nil {
		return err
	}

	return nil
}