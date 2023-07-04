package product

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"
)

func HandleProducts(writer http.ResponseWriter, request *http.Request) {

	log.Println("handleProducts")

	switch request.Method {
	case http.MethodGet:

		products, err := getProductList()
		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}

		parsedTemplate, err := template.ParseFiles("products.gohtml")
		if err != nil {
			log.Println(err)
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}
		err = parsedTemplate.Execute(writer, products)
		if err != nil {
			log.Fatal(err)
		}

	default:
		writer.WriteHeader(http.StatusNotImplemented)
		return
	}
}

func HandleProduct(writer http.ResponseWriter, request *http.Request) {

	urlPathSegments := strings.Split(request.URL.Path, "products/")
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
