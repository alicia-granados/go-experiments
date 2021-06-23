package main

import (
	"log"
	"os"
)

func main() {

	_, err := os.Open("sin-archivo.txt")
	if err != nil {
		//log.Panicln(err)// muestra tambien fecha y hora
		log.Panic(err)
	}
}

/* Panincln es equivalente a Println() seguido por una llmada a panic()*/
