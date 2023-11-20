package unitcontroller

import (
<<<<<<< HEAD
	"html/template"
	"net/http"
	"strconv"
	"time"
=======
	"html/template"
	"net/http"
>>>>>>> 6923ca4 (et

	"go-web/entities"
	"go-web/models/unitmodel"
	"go-web/models/unitmodel"
>>>>>>> 6923ca4 (et)
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

func Edit(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {

		temp, err := template.ParseFiles("views/unit/edit.html")
		if err != nil {
			panic(err)
		}

		idString := r.URL.Query().Get("id")
		id, err := strconv.Atoi(idString)
		if err != nil {
			panic(err)
		}

		units := unitmodel.Detail(id)

		data := map[string]any{
			"units": units,
		}

		temp.Execute(w, data)

	}

	if r.Method == "POST" {

		//menangkap parameter id dari Form Value
		idString := r.FormValue("id")
		id, err := strconv.Atoi(idString)
		if err != nil {
			panic(err)
		}

		var unit entities.Unit
		unit.Name = r.FormValue("name")
		unit.UpdatedAt = time.Now()

		if success := unitmodel.Update(id, unit); !success {
			panic(err)
		}

		http.Redirect(w, r, "/units", http.StatusSeeOther)
	}
}

func Delete(w http.ResponseWriter, r *http.Request) {

	// mengangkap parameter id dari URL
	idString := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		panic(err)
	}

	if err := unitmodel.Delete(id); err != nil {
		panic(err)
	}

	http.Redirect(w, r, "/units", http.StatusSeeOther)
}
