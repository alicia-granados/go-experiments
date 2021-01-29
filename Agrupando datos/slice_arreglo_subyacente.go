package main

import "fmt"

func main() {

	x := []int{1, 2, 3, 4, 5}
	fmt.Println(x)

	y := append(x, 6, 7, 8) // un nuevo areglo es asignado
	fmt.Println(y)

	z := append(x[:2], x[3:]...) // se utiliza el mismo arreglo subyacente
	fmt.Println(x)
	fmt.Println(z)

}
