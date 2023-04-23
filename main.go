package main

import (
	"go-web-native/config"
	categoriescontroller "go-web-native/controllers/categoriesController"
	homecontroller "go-web-native/controllers/homeController"
	productcontroller "go-web-native/controllers/productController"
	"log"
	"net/http"
)

func main() {
	config.ConnectDB()

	//homepage
	http.HandleFunc("/", homecontroller.Welcome)
	//categories
	http.HandleFunc("/categories", categoriescontroller.Index)
	http.HandleFunc("/categories/add", categoriescontroller.Add)
	http.HandleFunc("/categories/edit", categoriescontroller.Edit)
	http.HandleFunc("/categories/delete", categoriescontroller.Delete)
	//products
	http.HandleFunc("/products", productcontroller.Index)
	http.HandleFunc("/products/add", productcontroller.Add)
	http.HandleFunc("/products/edit", productcontroller.Edit)
	http.HandleFunc("/products/delete", productcontroller.Delete)

	log.Println("Server running on port 8080")
	http.ListenAndServe(":8080", nil)
}
