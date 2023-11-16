package pangkatcontroller

import (
	"go-web/entities"
	"go-web/models/pangkatmodel"
	"html/template"
	"net/http"
	"time"
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
