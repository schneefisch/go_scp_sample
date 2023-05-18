package database

import (
	"database/sql"
	_ "github.com/proullon/ramsql/driver"
)

var DbConn *sql.DB

// SetupDatabase will create a simple in-memory database an pre-fill it with some start-data
func SetupDatabase() error {

	batch := []string{
		`CREATE TABLE products (productId BIGSERIAL PRIMARY KEY, productName TEXT, price TEXT, quantityOnHand INT);`,
		`INSERT INTO products (productId, productName, price, quantityOnHand) VALUES (0, 'Edamer', '7.99', 15)`,
		`INSERT INTO products (productId, productName, price, quantityOnHand) VALUES (1, 'Gouda', '5.99', 22)`,
		`INSERT INTO products (productId, productName, price, quantityOnHand) VALUES (2, 'Mozzarella', '6.49', 12)`,
	}

	db, err := sql.Open("ramsql", "Products")
	if err != nil {
		return err
	}

	for _, query := range batch {
		_, err = db.Exec(query)
		if err != nil {
			return err
		}
	}

	DbConn = db

	return nil
}

// StopDatabase will destroy the Database-connection and with it the in-memory database
func StopDatabase() {
	_ = DbConn.Close()
}
