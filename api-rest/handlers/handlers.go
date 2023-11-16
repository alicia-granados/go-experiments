package handlers

import (
	"api-rest/db"
	"api-rest/models"
	"encoding/json"
	"fmt"
	"net/http"
)

func GetUsers(rw http.ResponseWriter, r *http.Request) {

	rw.Header().Set("Content-Type", "application/json")
	// rw.Header().Set("Content-Type", "text/xml")
	// en yaml no hay tipo de dato en especifico se va omitir la linea trece
	db.Connect()
	users := models.ListUser()
	db.Close()

	//transformar el objeto a atipo json
	output, err := json.Marshal(users)
	//transformar con xml 	->  output, err := xml.Marshal(users)
	// transformar yaml ->  output, err := yaml.Marshal(users)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Fprintln(rw, string(output))
}

func GetUser(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(rw, "Obtiene usuario")
}

func CreateUsers(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(rw, "Crea un usuario")
}

func UpdateUser(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(rw, "Actualiza usuario")
}

func DeleteUser(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(rw, "Elimina usuario")
}
