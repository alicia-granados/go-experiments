package main

import (
	"fmt"
	"net/http"
)

func main() {
	//colocar ruta donde se va a ejecutar la app
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hola mundo")
	})
	// crear servidor
	port := ":8080"
	fmt.Printf("servidor escuchando en http://localhost%s\n", port)
	http.ListenAndServe(port, nil)
}
