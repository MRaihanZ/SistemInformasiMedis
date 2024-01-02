package main

import (
	"go-web-native/config"
	"go-web-native/controllers/homecontroller"
	"go-web-native/controllers/productcontroller"
	"log"
	"net/http"
)

func main() {
	// 	Database connection
	config.ConnectDB()

	// Routes
	// Homepage
	http.HandleFunc("/", homecontroller.Welcome)

	// Products
	http.HandleFunc("/pasien", productcontroller.Index)
	http.HandleFunc("/pasien/add", productcontroller.Add)
	http.HandleFunc("/pasien/detail", productcontroller.Detail)
	http.HandleFunc("/pasien/edit", productcontroller.Edit)
	http.HandleFunc("/pasien/delete", productcontroller.Delete)

	// Info
	log.Println("Server running on port: 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
