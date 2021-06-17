package main

import (
	"fmt"
)

func main() {

	a := 42
	fmt.Println(a)
	fmt.Println(&a) //dirección de memoria

	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", &a)

	var b *int = &a
	fmt.Println(b)
	fmt.Println(*b) // obrtener valor

	*b = 43
	fmt.Println(a)
}
