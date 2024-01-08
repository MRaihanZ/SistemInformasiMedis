package printcontroller

import (
	"go-web-native/models/printmodel"
	"html/template"
	"net/http"
	"strconv"
)

func Print(w http.ResponseWriter, r *http.Request) {
	idString := r.URL.Query().Get("id")

	id, err := strconv.Atoi(idString)
	if err != nil {
		panic(err)
	}

	records := printmodel.GetAll(id)
	data := map[string]any{
		"records": records,
	}

	temp, err := template.ParseFiles("views/pasien/print.html")
	if err != nil {
		panic(err)
	}

	temp.Execute(w, data)
}