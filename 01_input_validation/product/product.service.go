package product

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

const path = "products"

func SetupRoutes() {
	productsHandler := http.HandlerFunc(handleProducts)
	productHandler := http.HandlerFunc(handleProduct)
	http.Handle(fmt.Sprintf("%s", path), productsHandler)
	http.Handle(fmt.Sprintf("%s/", path), productHandler)
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

func handleProduct(writer http.ResponseWriter, request *http.Request) {

	urlPathSegments := strings.Split(request.URL.Path, fmt.Sprintf("%s/", path))
	if len(urlPathSegments[1:]) > 1 {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	productId := urlPathSegments[len(urlPathSegments)-1]

	switch request.Method {
	case http.MethodGet:

		product, err := getProductById(productId)
		if err != nil {
			log.Println(err)
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}
		if product == nil {
			log.Println(fmt.Sprintf("No product found for id %s", productId))
			writer.WriteHeader(http.StatusNotFound)
			return
		}
		productJson, err := json.Marshal(product)
		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}
		writer.WriteHeader(http.StatusOK)
		_, err = writer.Write(productJson)
		if err != nil {
			log.Fatal(err)
		}
	default:
		writer.WriteHeader(http.StatusNotImplemented)
	}

}
