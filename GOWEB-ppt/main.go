package main

import (
	"GOWEB-ppt/handlers"
	"fmt"
	"log"
	"net/http"
)

func main() {

	//Crear enrutador
	router := http.NewServeMux()

	// configurar rutas
	router.HandleFunc("/", handlers.Index)
	router.HandleFunc("/new", handlers.NewGame)
	router.HandleFunc("/game", handlers.Game)
	router.HandleFunc("/play", handlers.Play)
	router.HandleFunc("/about", handlers.About)

	//colocar ruta donde se va a ejecutar la app
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hola mundo")
	})
	// crear servidor
	port := ":8080"
	log.Printf("servidor escuchando en http://localhost%s\n", port)
	log.Fatal(http.ListenAndServe(port, router))
}
