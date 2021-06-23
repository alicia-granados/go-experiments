package main

import (
	"log"
	"os"
)

func main() {
	_, err := os.Open("sin-archivo.txt")

	if err != nil {
		log.Println("Ocurrió un error", err)
	}
	// con log muestra tambien la hora y la fecha
}
