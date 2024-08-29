package main

import (
	"log"
	"profiles_go/db"
	"profiles_go/routes"
)

func main() {
	// Initialize the database
	err := db.InitDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Start the server
	routes.Routes()

}
