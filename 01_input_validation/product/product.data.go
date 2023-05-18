package product

import (
	"github.com/schneefisch/go_scp_sample/01_input_validation/database"
	"log"
)

//goland:noinspection SqlNoDataSourceInspection
func getProductList() ([]Product, error) {
	results, err := database.DbConn.Query(`SELECT 
		productId, 
		productName, 
		price, 
		quantityOnHand 
	FROM products`)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer results.Close()

	products := make([]Product, 0)
	for results.Next() {
		var product Product
		err := results.Scan(&product.ProductId,
			&product.ProductName,
			&product.Price,
			&product.QuantityOnHand)
		if err != nil {
			log.Println(err)
			return nil, err
		}

		products = append(products, product)
	}
	return products, nil
}
