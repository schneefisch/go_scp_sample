package product

import (
	"github.com/schneefisch/go_scp_sample/database"
	"log"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	// setup test-data
	err := database.SetupDatabase()
	if err != nil {
		log.Println("Error setting up the database")
		log.Fatal(err)
	}

	// run tests
	code := m.Run()

	// cleanup
	database.StopDatabase()

	os.Exit(code)
}
