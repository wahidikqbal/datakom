package stockcontroller

import (
	"html/template"
	"net/http"
	"strconv"
	"time"

	"go-web/entities"
	"go-web/models/stockmodel"
)

func Index(w http.ResponseWriter, r *http.Request) {
	temp, err := template.ParseFiles("views/stock/index.html")
	if err != nil {
		panic(err)
	}

	stocks := stockmodel.GetAll()

	data := map[string]any{
		"stocks": stocks,
	}

	temp.Execute(w, data)

}

func Add(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		temp, err := template.ParseFiles("views/stock/create.html")
		if err != nil {
			panic(err)
		}

		temp.Execute(w, nil)
	}

	if r.Method == "POST" {
		var stock entities.Stock
		stock.Name = r.FormValue("name")
		stock.CreatedAt = time.Now()
		stock.UpdatedAt = time.Now()

		if success := stockmodel.Create(stock); !success {
			temp, err := template.ParseFiles("views/stock/create.html")
			if err != nil {
				panic(err)
			}

			temp.Execute(w, nil)
		}

		http.Redirect(w, r, "/stocks", http.StatusSeeOther)
	}
}

func Edit(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		temp, err := template.ParseFiles("views/stock/edit.html")
		if err != nil {
			panic(err)
		}

		//menangakap parameter id dari URL
		idString := r.URL.Query().Get("id")
		id, err := strconv.Atoi(idString)
		if err != nil {
			panic(err)
		}

		stocks := stockmodel.Detail(id)

		data := map[string]any{
			"stocks": stocks,
		}

		temp.Execute(w, data)
	}

	if r.Method == "POST" {

		idString := r.FormValue("id")
		id, err := strconv.Atoi(idString)
		if err != nil {
			panic(err)
		}

		var stock entities.Stock
		stock.Name = r.FormValue("name")
		stock.UpdatedAt = time.Now()

		if success := stockmodel.Update(id, stock); !success {
			panic(err)
		}

		http.Redirect(w, r, "/stocks", http.StatusSeeOther)
	}

}

func Delete(w http.ResponseWriter, r *http.Request) {
	idString := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		panic(err)
	}

	if err := stockmodel.Delete(id); err != nil {
		panic(err)
	}

	http.Redirect(w, r, "/stocks", http.StatusSeeOther)
}
