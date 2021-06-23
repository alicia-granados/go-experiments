package main

import (
	"fmt"
	"log"
	"os"
)

func main() {

	defer foo()

	_, err := os.Open("sin-archivo.txt")
	if err != nil {
		log.Fatalln(err) // Todo muere
	}
}

func foo() {
	fmt.Println("Cuando os.Exit() es llamada, las funciones diferidas no corren")
}

/* las funciones de fatal llaman a os.Exit() después de escribir el mensaje del log

Fatalln es equivalente a Println() seguido por una llamada a os. Exit(1)*/
