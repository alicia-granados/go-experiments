package main

import (
	"fmt"
	"os"
)

func main() {
	_, err := os.Open("sin-archivo.txt")

	if err != nil {
		fmt.Println("Ocurrió un error", err)
	}
}
