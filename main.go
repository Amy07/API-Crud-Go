package main

import (
	"crud-go/controllers"
	"crud-go/database"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func RegisterProductRoutes(router *mux.Router) {
	router.HandleFunc("/getall", controllers.GetProducts).Methods("GET")
	router.HandleFunc("/get/{id}", controllers.GetProductById).Methods("GET")
	router.HandleFunc("/create", controllers.CreateProduct).Methods("POST")
	router.HandleFunc("/update/{id}", controllers.UpdateProduct).Methods("PUT")
	router.HandleFunc("/delete/{id}", controllers.DeleteProduct).Methods("DELETE")
}

func main() {

	LoadAppConfig()

	database.Connect(AppConfig.ConnectionString)
	database.Migrate()

	router := mux.NewRouter().StrictSlash(true)
	RegisterProductRoutes(router)

	log.Fatal(http.ListenAndServe(":8090", router))
}
