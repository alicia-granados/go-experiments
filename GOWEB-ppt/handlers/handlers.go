package handlers

import (
	"fmt"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Página de inicio")
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
