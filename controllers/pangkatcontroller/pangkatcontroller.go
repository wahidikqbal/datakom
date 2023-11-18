package pangkatcontroller

import (
	"html/template"
	"net/http"
	"strconv"
	"time"

	"go-web/entities"
	"go-web/models/pangkatmodel"
)

func Index(w http.ResponseWriter, r *http.Request) {
	pangkats := pangkatmodel.GetAll()

	data := map[string]any{
		"pangkats": pangkats,
	}

	temp, err := template.ParseFiles("views/pangkat/index.html")
	if err != nil {
		panic(err)
	}

	temp.Execute(w, data)
}

func Add(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		temp, err := template.ParseFiles("views/pangkat/create.html")
		if err != nil {
			panic(err)
		}

		temp.Execute(w, nil)
	}

	if r.Method == "POST" {
		var pangkat entities.Pangkat
		pangkat.Name = r.FormValue("name")
		pangkat.CreatedAt = time.Now()
		pangkat.UpdatedAt = time.Now()

		if success := pangkatmodel.Create(pangkat); !success {
			temp, err := template.ParseFiles("views/pangkat/create.html")
			if err != nil {
				panic(err)
			}

			temp.Execute(w, nil)
		}

		http.Redirect(w, r, "/pangkats", http.StatusSeeOther)
	}

}

func Edit(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		temp, err := template.ParseFiles("views/pangkat/edit.html")
		if err != nil {
			panic(err)
		}

		//menangkap paremeter id dari reques
		idStirng := r.URL.Query().Get("id")
		id, err := strconv.Atoi(idStirng)
		if err != nil {
			panic(err)
		}

		pangkat := pangkatmodel.Detail(id)

		data := map[string]any{
			"pangkat": pangkat,
		}

		temp.Execute(w, data)

	}

	if r.Method == "POST" {
		var pangkat entities.Pangkat

		//menangkap parameter id dari form value
		idString := r.FormValue("id")
		id, err := strconv.Atoi(idString)
		if err != nil {
			panic(err)
		}

		pangkat.Name = r.FormValue("name")
		pangkat.UpdatedAt = time.Now()

		if success := pangkatmodel.Update(id, pangkat); !success {
			panic(err)
		}

		http.Redirect(w, r, "/pangkats", http.StatusSeeOther)

	}
}
