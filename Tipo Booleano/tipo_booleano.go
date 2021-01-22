package main

import "fmt"

var x bool

func main() {

	fmt.Println(x)
	x = true

	fmt.Println(x)

	z := 7
	y := 42

	fmt.Println(z == y)
	fmt.Println(z <= y)
	fmt.Println(z >= y)
	fmt.Println(z != y)

}
