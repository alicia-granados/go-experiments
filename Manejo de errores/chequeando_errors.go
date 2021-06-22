package main

import (
	"fmt"
)

func main() {

	var respuesta1, respuesta2, respuesta3 string

	fmt.Print("Nombre: ")
	_, err := fmt.Scan(&respuesta1)
	if err != nil {
		panic(err)
	}

	fmt.Print("Comida Favorita: ")
	_, err = fmt.Scan(&respuesta2)
	if err != nil {
		panic(err)
	}

	fmt.Print("Deporte Favorito: ")
	_, err = fmt.Scan(&respuesta3)
	if err != nil {
		panic(err)
	}

	fmt.Println(respuesta1, respuesta2, respuesta3)
}
