package main

import (
	"go-web-native/config"
	"go-web-native/controllers/homecontroller"
	"go-web-native/controllers/recordcontroller"
	"log"
	"net/http"
)

func main() {
	// 	Database connection
	config.ConnectDB()

	// Routes
	// Homepage
	http.HandleFunc("/", homecontroller.Welcome)

	// Pasien
	http.HandleFunc("/pasien", recordcontroller.Index)
	http.HandleFunc("/pasien/add", recordcontroller.Add)
	http.HandleFunc("/pasien/detail", recordcontroller.Detail)
	http.HandleFunc("/pasien/edit", recordcontroller.Edit)
	http.HandleFunc("/pasien/delete", recordcontroller.Delete)

	// Info
	log.Println("Server running on port: 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
