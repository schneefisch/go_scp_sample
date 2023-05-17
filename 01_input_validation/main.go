package main

import (
	"github.com/schneefisch/go_scp_sample/01_input_validation/product"
	"log"
	"net/http"
)

func main() {
	// ToDo: Setup in-memory database

	// Setup Routes
	product.SetupRoutes()

	// Start Server
	log.Fatal(http.ListenAndServe(":5000", nil))
}
