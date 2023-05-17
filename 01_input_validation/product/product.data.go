package product

func getProductList() ([]Product, error) {
	return []Product{
		{
			ProductId:      1,
			ProductName:    "Edamer Käse",
			Price:          "7.99",
			QuantityOnHand: 5,
		},
		{
			ProductId:      2,
			ProductName:    "Gauda Käse",
			Price:          "5,99",
			QuantityOnHand: 3,
		},
	}, nil
}
