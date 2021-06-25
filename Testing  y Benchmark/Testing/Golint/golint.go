package main

import (
	"fmt"
)

func main() {
	fmt.Println("La suma de 2 + 4 es: ", miSuma(2, 4))
	fmt.Println("La suma de 1 + 4 es: ", miSuma(1, 4))
	fmt.Println("La suma de 6 + 7 es: ", miSuma(6, 7))
}

func miSuma(xi ...int) int {
	var sum int
	for _, v := range xi {
		sum += v
	}
	return sum
}

//gofmt golint.go Reformatea el codigo
//go vet ./... muetsra los archivos que tienes errores
//golint ./... ayuda a mostrar los archivos que tienes errores de estilo
