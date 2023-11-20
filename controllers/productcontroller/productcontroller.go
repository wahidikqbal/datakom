package productcontroller

import (
	"net/http"
	"text/template"

	"go-web/models/categorymodel"
	"go-web/models/pangkatmodel"
	"go-web/models/productmodel"
	"go-web/models/stockmodel"
	"go-web/models/unitmodel"
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
		units := unitmodel.GetAll()
		stocks := stockmodel.GetAll()

		data := map[string]any{
			"categories": categories,
			"pangkats":   pangkats,
			"units":      units,
			"stocks":     stocks,
		}

		temp.Execute(w, data)
	}

}
