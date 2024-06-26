package main

import (
	"apirest-gorm/handlers"
	"apirest-gorm/models"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	models.MigrarUser()

	//rutas
	mux := mux.NewRouter()
	//Endpoint
	mux.HandleFunc("/api/user/", handlers.GetUsers).Methods("GET")
	mux.HandleFunc("/api/user/{id:[0-9]+}", handlers.GetUser).Methods("GET")
	mux.HandleFunc("/api/user/", handlers.CreateUsers).Methods("POST")
	mux.HandleFunc("/api/user/{id:[0-9]+}", handlers.UpdateUser).Methods("PUT")
	mux.HandleFunc("/api/user/{id:[0-9]+}", handlers.DeleteUser).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":3001", mux))
}
