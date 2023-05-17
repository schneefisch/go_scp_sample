package product

type Product struct {
	ProductId      int    `json:"product_id"`
	ProductName    string `json:"product_name"`
	Price          string `json:"price"`
	QuantityOnHand int    `json:"quantity_on_hand"`
}
