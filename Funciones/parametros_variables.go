package main

import "fmt"

func main() {
	x := suma(2, 3, 4, 5, 6, 7, 8, 9, 10)
	fmt.Println("EL valor total almacenado e la variable es ", x)
}

//func (r receptor) identificador(parámetros ) (retorno(s)){código}
func suma(x ...int) int {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	suma := 0
	for i, v := range x {
		suma += v
		fmt.Println("EL valor en el índice ", i, " suma ", v, "al total, quedando ", suma)
	}
	fmt.Println("El toral es", suma)
	return suma
}
