package main

import (
	"log"
	"net/http"

	"go-web/config"
	"go-web/controllers/categorycontroller"
	"go-web/controllers/homecontroller"
	"go-web/controllers/pangkatcontroller"
	"go-web/controllers/productcontroller"
	"go-web/controllers/stockcontroller"
	"go-web/controllers/unitcontroller"
)

func main() {

	//1. Homepage
	http.HandleFunc("/", homecontroller.Welcome)

	//2. Category
	http.HandleFunc("/categories", categorycontroller.Index)
	http.HandleFunc("/categories/add", categorycontroller.Add)
	http.HandleFunc("/categories/edit", categorycontroller.Edit)
	http.HandleFunc("/categories/delete", categorycontroller.Delete)

	//3. Product
	http.HandleFunc("/products", productcontroller.Index)
	http.HandleFunc("/products/add", productcontroller.Add)

	//3. Pangkat
	http.HandleFunc("/pangkats", pangkatcontroller.Index)
	http.HandleFunc("/pangkats/add", pangkatcontroller.Add)
	http.HandleFunc("/pangkats/edit", pangkatcontroller.Edit)
	http.HandleFunc("/pangkats/delete", pangkatcontroller.Delete)

	//4. Unit
	http.HandleFunc("/units", unitcontroller.Index)
	http.HandleFunc("/units/add", unitcontroller.Add)
	http.HandleFunc("/units/edit", unitcontroller.Edit)
	http.HandleFunc("/units/delete", unitcontroller.Delete)

	//5. Stock
	http.HandleFunc("/stocks", stockcontroller.Index)
	http.HandleFunc("/stocks/add", stockcontroller.Add)
	http.HandleFunc("/stocks/edit", stockcontroller.Edit)
	http.HandleFunc("/stocks/delete", stockcontroller.Delete)

	//connect database
	config.ConnectDB()

	log.Println("server running on port 8080")
	http.ListenAndServe(":8080", nil)
}
