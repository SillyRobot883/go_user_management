package utils

import (
	"github.com/joho/godotenv"
	"log"
)

func LoadEnv() error {
	// load .env file
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found")
	}

	return err
}
