package handlers

import (
	"fmt"
	"html/template"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	tpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, "error al analizar plantillas", http.StatusInternalServerError)
		return
	}
	data := struct {
		Title   string
		Message string
	}{
		Title:   "pagina de inicio",
		Message: "Bienvenido a piedra, papel o tijera",
	}

	err = tpl.Execute(w, data)
	if err != nil {
		http.Error(w, "error al renderizar plantilla", http.StatusInternalServerError)
	}
}

func NewGame(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "crear nuevo juego")
}

func Game(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Juego")
}

func Play(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "jugar")
}

func About(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Acerca de ")
}