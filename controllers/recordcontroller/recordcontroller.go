package recordcontroller

import (
	"go-web-native/entities"
	"go-web-native/models/categorymodel"
	"go-web-native/models/doctormodel"
	"go-web-native/models/medicinemodel"
	"go-web-native/models/pasienmodel"
	"go-web-native/models/recordmodel"
	"net/http"
	"strconv"
	"text/template"
	"time"
)

func Index(w http.ResponseWriter, r *http.Request) {
	idString := r.URL.Query().Get("id")

	var isIdSet bool
	id, err := strconv.Atoi(idString)
	if err != nil {
		id = 0
		isIdSet = false
	} else {
		isIdSet = true
	}

	records := recordmodel.Getall(isIdSet, id)
	data := map[string]any{
		"records": records,
	}

	temp, err := template.ParseFiles("views/pasien/index.html")
	if err != nil {
		panic(err)
	}

	temp.Execute(w, data)
}

func Add(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		temp, err := template.ParseFiles("views/pasien/create.html")
		if err != nil {
			panic(err)
		}

		pasien := pasienmodel.GetAll()
		categories := categorymodel.GetAll()
		doctor := doctormodel.GetAll()
		medicine := medicinemodel.GetAll()
		data := map[string]any{
			"pasien": pasien,
			"categories": categories,
			"doctor": doctor,
			"medicine": medicine,
		}

		temp.Execute(w, data)
	}

	if r.Method == "POST" {
		var record entities.Record

		idPasien, err := strconv.Atoi(r.FormValue("id_pasien"))
		if err != nil {
			panic(err)
		}

		idCategory, err := strconv.Atoi(r.FormValue("id_category"))
		if err != nil {
			panic(err)
		}

		idDoctor, err := strconv.Atoi(r.FormValue("id_doctor"))
		if err != nil {
			panic(err)
		}

		diagnose := r.FormValue("diagnose")

		description := r.FormValue("description")

		idMedicine, err := strconv.Atoi(r.FormValue("id_medicine"))
		if err != nil {
			panic(err)
		}

		record.Pasien.Id = uint(idPasien)
		record.Category.Id = uint(idCategory)
		record.Doctor.Id = uint(idDoctor)
		record.Diagnose = diagnose
		record.Description = description
		record.Medicine.Id = uint(idMedicine)
		record.CreatedAt = time.Now()
		record.UpdatedAt = time.Now()

		if ok := recordmodel.Create(record); !ok {
			http.Redirect(w, r, "/pasien", http.StatusTemporaryRedirect)
			return
		}

		http.Redirect(w, r, "/pasien", http.StatusSeeOther)
	}
}

func Detail(w http.ResponseWriter, r *http.Request) {
	idString := r.URL.Query().Get("id")

	id, err := strconv.Atoi(idString)
	if err != nil {
		panic(err)
	}

	record := recordmodel.Detail(id)
	data := map[string]any{
		"record": record,
	}

	temp, err := template.ParseFiles("views/pasien/detail.html")
	if err != nil {
		panic(err)
	}

	temp.Execute(w, data)
}

func Edit(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		temp, err := template.ParseFiles("views/pasien/edit.html")
		if err != nil {
			panic(err)
		}

		idString := r.URL.Query().Get("id")
		id, err := strconv.Atoi(idString)
		if err != nil {
			panic(err)
		}

		record := recordmodel.Detail(id)
		pasien := pasienmodel.GetAll()
		categories := categorymodel.GetAll()
		doctor := doctormodel.GetAll()
		medicine := medicinemodel.GetAll()
		data := map[string]any{
			"record": record,
			"pasien": pasien,
			"categories": categories,
			"doctor": doctor,
			"medicine": medicine,
		}

		temp.Execute(w, data)
	}

	if r.Method == "POST" {
		var record entities.Record

		id, err := strconv.Atoi(r.FormValue("id"))
		if err != nil {
			panic(err)
		}

		idPasien, err := strconv.Atoi(r.FormValue("id_pasien"))
		if err != nil {
			panic(err)
		}

		idCategory, err := strconv.Atoi(r.FormValue("id_category"))
		if err != nil {
			panic(err)
		}

		idDoctor, err := strconv.Atoi(r.FormValue("id_doctor"))
		if err != nil {
			panic(err)
		}

		diagnose := r.FormValue("diagnose")

		description := r.FormValue("description")

		idMedicine, err := strconv.Atoi(r.FormValue("id_medicine"))
		if err != nil {
			panic(err)
		}

		record.Id = uint(id)
		record.Pasien.Id = uint(idPasien)
		record.Category.Id = uint(idCategory)
		record.Doctor.Id = uint(idDoctor)
		record.Diagnose = diagnose
		record.Description = description
		record.Medicine.Id = uint(idMedicine)
		record.UpdatedAt = time.Now()

		if ok := recordmodel.Update(record); !ok {
			http.Redirect(w, r, "/pasien", http.StatusTemporaryRedirect)
			return
		}

		http.Redirect(w, r, "/pasien", http.StatusSeeOther)
	}
}

func Delete(w http.ResponseWriter, r *http.Request) {
	idString := r.URL.Query().Get("id")

	id, err := strconv.Atoi(idString)
	if err != nil {
		panic(err)
	}

	if err := recordmodel.Delete(id); err != nil {
		panic(err)
	}

	http.Redirect(w, r, "/pasien", http.StatusSeeOther)
}