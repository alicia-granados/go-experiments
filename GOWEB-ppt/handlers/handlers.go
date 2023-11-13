package handlers

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

const (
	templateDir  = "templates/"
	templateBase = templateDir + "base.html"
)

func Index(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "index.html", nil)
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

func renderTemplate(w http.ResponseWriter, page string, data any) {

	tpl := template.Must(template.ParseFiles(templateBase, templateDir+page))

	err := tpl.ExecuteTemplate(w, "base", data)
	if err != nil {
		http.Error(w, "error al renderizar plantilla", http.StatusInternalServerError)
		log.Println(err)
		return
	}
}
