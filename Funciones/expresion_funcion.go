package main

import (
	"fmt"
)

func main() {

	f := func() {
		fmt.Println("Mi primera expresión función")
	}
	f()

	g := func(y int) {
		fmt.Println("El año en que se descubrió América fué :", y)
	}
	g(1492)
}
