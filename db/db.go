package db

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
	"os"
	"profiles_go/utils"
)

var DB *sql.DB

func InitDB() error {
	var err error

	utils.LoadEnv()

	connStr := os.Getenv("DATABASE_URL")
	if connStr == "" {
		log.Fatal("DATABASE_URL is not set in the environment variables")
	}

	// Get the environment variables
	DB, err = sql.Open("postgres", string(connStr))
	if err != nil {
		return err
	}

	if err = DB.Ping(); err != nil {
		return err
	}

	// Create tables
	if err = createTables(); err != nil {
		return err
	}

	return nil
}

func Close() {
	if DB != nil {
		DB.Close()
	}
}

func createTables() error {
	queries := []string{
		`CREATE TABLE IF NOT EXISTS users (
			id SERIAL PRIMARY KEY,
			username VARCHAR(50) UNIQUE NOT NULL,
			email VARCHAR(255) UNIQUE NOT NULL,
			password_hash VARCHAR(255) NOT NULL,
			created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
			updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
		)`,
		`CREATE TABLE IF NOT EXISTS profiles (
			id SERIAL PRIMARY KEY,
			user_id INT REFERENCES users(id) ON DELETE CASCADE,
			avatar TEXT,
			bio TEXT,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
			);`,
	}

	for _, query := range queries {
		_, err := DB.Exec(query)
		if err != nil {
			return err
		}
	}

	return nil
}
