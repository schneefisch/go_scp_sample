package database

import (
	"database/sql"
	_ "github.com/proullon/ramsql/driver"
)

var DbConn *sql.DB

// SetupDatabase will create a simple in-memory database an pre-fill it with some start-data
//
//goland:noinspection SqlNoDataSourceInspection
func SetupDatabase() error {

	batch := []string{
		`CREATE TABLE products (productId BIGSERIAL PRIMARY KEY, productName TEXT, price TEXT, quantityOnHand INT);`,
		`INSERT INTO products (productName, price, quantityOnHand) VALUES ('Edamer', '7.99', 15)`,
		`INSERT INTO products (productName, price, quantityOnHand) VALUES ('Gouda', '5.99', 22)`,
		`INSERT INTO products (productName, price, quantityOnHand) VALUES ('Mozzarella', '6.49', 12)`,
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
