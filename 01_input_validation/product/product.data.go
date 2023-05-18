package product

import (
	"context"
	"fmt"
	"github.com/schneefisch/go_scp_sample/01_input_validation/database"
	"log"
)

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

func getProductById(productId string) (*Product, error) {
	ctx := context.Background()
	query := fmt.Sprintf(`SELECT productId, productName, price, quantityOnHand FROM products WHERE productId = %s`, productId)
	result, err := database.DbConn.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer result.Close()

	var product Product
	for result.Next() {
		err := result.Scan(&product.ProductId,
			&product.ProductName,
			&product.Price,
			&product.QuantityOnHand)
		if err != nil {
			log.Println(err)
			return nil, err
		}
	}
	return &product, nil
}
