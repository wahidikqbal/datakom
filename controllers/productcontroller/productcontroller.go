package productcontroller

import (
	"net/http"
	"text/template"
	"time"

	"go-web/entities"
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

	if r.Method == "POST" {

		var product entities.Product

		// pangkat, err := strconv.Atoi(r.FormValue("pangkat_id"))
		// if err != nil {
		// 	panic(err)
		// }

		// unit, err := strconv.Atoi(r.FormValue("unit_id"))
		// if err != nil {
		// 	panic(err)
		// }

		// category, err := strconv.Atoi(r.FormValue("category_id"))
		// if err != nil {
		// 	panic(err)
		// }

		// stock, err := strconv.Atoi(r.FormValue("stock_id"))
		// if err != nil {
		// 	panic(err)
		// }

		product.Name = r.FormValue("name")
		// product.Pangkat.Id = uint(pangkat)
		// product.Nrp = r.FormValue("nrp")
		// product.Unit.Id = uint(unit)
		// product.Category.Id = uint(category)
		// product.Serialnumber = r.FormValue("serialnumber")
		// product.Stock.Id = uint(stock)
		product.CreatedAt = time.Now()
		product.UpdatedAt = time.Now()

		if success := productmodel.Create(product); !success {
			http.Redirect(w, r, r.Header.Get("Referer"), http.StatusTemporaryRedirect)
		}

		http.Redirect(w, r, "/products", http.StatusSeeOther)
	}

}
