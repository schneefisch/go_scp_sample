package main

import (
	"github.com/schneefisch/go_scp_sample/app"
	"github.com/schneefisch/go_scp_sample/database"
	"log"
	"net/http"
)

func main() {
	// Setup in-memory database
	err := database.SetupDatabase()
	if err != nil {
		log.Fatal("Error setting up the database connection")
	}
	defer database.StopDatabase()

	// Setup Routes
	app.SetupRoutes()

	// Start Server
	log.Fatal(http.ListenAndServe(":8080", nil))
}
