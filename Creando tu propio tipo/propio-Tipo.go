package main

import "fmt"

var a int

type dinero int

var b dinero

func main() {
	a = 42
	fmt.Println(a)
	fmt.Printf("%T\n", a)

	b = 1000
	fmt.Println(b)
	fmt.Printf("%T\n", b)

	/* no se puede asignar mostrará un error
	a = b
	dado que a es de tipo int y b es de tipo entero*/
}
