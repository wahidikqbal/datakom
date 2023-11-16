package pangkatcontroller

import (
	"go-web/models/pangkatmodel"
	"html/template"
	"net/http"
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
