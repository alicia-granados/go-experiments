package handlers

import (
	"fmt"
	"net/http"
)

func GetUsers(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(rw, "lista de todos los usuarios")
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
