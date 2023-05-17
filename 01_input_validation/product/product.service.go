package product

import (
	"encoding/json"
	"log"
	"net/http"
)

func SetupRoutes() {
	productsHandler := http.HandlerFunc(handleProducts)
	http.Handle("products", productsHandler)
}

func handleProducts(writer http.ResponseWriter, request *http.Request) {

	switch request.Method {
	case http.MethodGet:

		products, err := getProductList()
		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}
		productsJson, err := json.Marshal(products)
		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}
		_, err = writer.Write(productsJson)
		if err != nil {
			log.Fatal(err)
			return
		}

	default:
		writer.WriteHeader(http.StatusNotImplemented)
		return
	}

}
