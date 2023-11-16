package productcontroller

import (
	"go-web/models/categorymodel"
	"go-web/models/pangkatmodel"
	"go-web/models/productmodel"
	"net/http"
	"text/template"
)

func Index(w http.ResponseWriter, r *http.Request) {
	products := productmodel.GetAll()

	data := map[string]any{
		"products": products,
	}

	temp, err := template.ParseFiles("views/product/index.html")
	if err != nil {
		panic(err)
	}

	temp.Execute(w, data)

}

func Add(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		temp, err := template.ParseFiles("views/product/create.html")
		if err != nil {
			panic(err)
		}

		categories := categorymodel.GetAll()
		pangkats := pangkatmodel.GetAll()

		data := map[string]any{
			"categories": categories,
			"pangkats":   pangkats,
		}

		temp.Execute(w, data)
	}

}
