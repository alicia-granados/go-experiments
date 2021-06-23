package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	f, err := os.Create("log.txt")

	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()
	log.SetOutput(f)

	f2, err := os.Open("sin-archivo.txt")
	if err != nil {
		log.Println("Un error ocurrió", err)
	}
	defer f2.Close()

	fmt.Println("Revisa el archivo log.txt en el directorio")
}
