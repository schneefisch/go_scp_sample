package app

import (
	"fmt"
	"github.com/schneefisch/go_scp_sample/product"
	"html/template"
	"log"
	"net/http"
)

const ProductPath = "products"

func SetupRoutes() {
	welcomeHandler := http.HandlerFunc(handleWelcome)
	productsHandler := http.HandlerFunc(product.HandleProducts)
	productHandler := http.HandlerFunc(product.HandleProduct)
	http.Handle(fmt.Sprintf("/%s/", ProductPath), productHandler)
	http.Handle(fmt.Sprintf("/%s", ProductPath), productsHandler)
	http.Handle("/", welcomeHandler)
}

func handleWelcome(writer http.ResponseWriter, request *http.Request) {
	switch request.Method {
	case http.MethodGet:

		parsedTemplate, err := template.ParseFiles("template/index.gohtml")
		if err != nil {
			log.Println(err)
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}
		err = parsedTemplate.Execute(writer, nil)
		if err != nil {
			log.Fatal(err)
		}
	}
}
