package categorycontroller

import (
	"html/template"
	"net/http"
	"strconv"
	"time"

	"go-web/entities"
	"go-web/models/categorymodel"
)

func Index(w http.ResponseWriter, r *http.Request) {
	categories := categorymodel.GetAll()

	data := map[string]any{
		"categories": categories,
	}

	temp, err := template.ParseFiles("views/category/index.html")
	if err != nil {
		panic(err)
	}

	temp.Execute(w, data)
}

func Add(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		temp, err := template.ParseFiles("views/category/create.html")
		if err != nil {
			panic(err)
		}

		temp.Execute(w, nil)
	}

	if r.Method == "POST" {
		var category entities.Category
		category.Name = r.FormValue("name")
		category.CreatedAt = time.Now()
		category.UpdatedAt = time.Now()

		if success := categorymodel.Create(category); !success {
			temp, _ := template.ParseFiles("/views/category/create.html")
			temp.Execute(w, nil)
		}

		http.Redirect(w, r, "/categories", http.StatusSeeOther)
	}
}

func Edit(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		temp, err := template.ParseFiles("views/category/edit.html")
		if err != nil {
			panic(err)
		}

		// Menangkap parameter id dari request
		idstring := r.URL.Query().Get("id")
		id, err := strconv.Atoi(idstring)
		if err != nil {
			panic(err)
		}

		category := categorymodel.Detail(id)

		data := map[string]any{
			"category": category,
		}

		temp.Execute(w, data)
	}

	if r.Method == "POST" {
		var category entities.Category

		//menangkap paramater id dari request form value
		idstring := r.FormValue("id")
		id, err := strconv.Atoi(idstring)
		if err != nil {
			panic(err)
		}

		category.Name = r.FormValue("name")
		category.UpdatedAt = time.Now()

		if success := categorymodel.Update(id, category); !success {
			http.Redirect(w, r, r.Header.Get("Referer"), http.StatusSeeOther)
			return
		}

		http.Redirect(w, r, "/categories", http.StatusSeeOther)

	}

}

func Delete(w http.ResponseWriter, r *http.Request) {

	//menangkap paramater id dari request form value
	idstring := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idstring)
	if err != nil {
		panic(err)
	}

	if err := categorymodel.Delete(id); err != nil {
		panic(err)
	}

	http.Redirect(w, r, "/categories", http.StatusSeeOther)

}

func Delete2(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		temp, err := template.ParseFiles("views/category/delete.html")
		if err != nil {
			panic(err)
		}

		// Menangkap parameter id dari request
		idstring := r.URL.Query().Get("id")
		id, err := strconv.Atoi(idstring)
		if err != nil {
			panic(err)
		}

		category := categorymodel.Detail(id)

		data := map[string]any{
			"category": category,
		}

		temp.Execute(w, data)
	}

	if r.Method == "DELETE" {
		//menangkap paramater id dari request form value
		idstring := r.FormValue("id")
		id, err := strconv.Atoi(idstring)
		if err != nil {
			panic(err)
		}

		if success := categorymodel.Delete2(id); !success {
			http.Redirect(w, r, r.Header.Get("Referer"), http.StatusSeeOther)
			return
		}

		http.Redirect(w, r, "/categories", http.StatusSeeOther)

	}

}
