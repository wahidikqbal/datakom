package unitcontroller

import (
	"go-web/entities"
	"go-web/models/unitmodel"
	"html/template"
	"net/http"
	"time"
)

func Index(w http.ResponseWriter, r *http.Request) {
	units := unitmodel.GetAll()

	data := map[string]any{
		"units": units,
	}

	temp, err := template.ParseFiles("views/unit/index.html")
	if err != nil {
		panic(err)
	}

	temp.Execute(w, data)
}

func Add(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		temp, err := template.ParseFiles("views/unit/create.html")
		if err != nil {
			panic(err)
		}

		temp.Execute(w, nil)
	}

	if r.Method == "POST" {
		var unit entities.Unit
		unit.Name = r.FormValue("name")
		unit.CreatedAt = time.Now()
		unit.UpdatedAt = time.Now()

		if success := unitmodel.Create(unit); !success {
			temp, err := template.ParseFiles("views/unit/create.html")
			if err != nil {
				panic(err)
			}

			temp.Execute(w, nil)
		}

		http.Redirect(w, r, "/units", http.StatusSeeOther)
	}

}
