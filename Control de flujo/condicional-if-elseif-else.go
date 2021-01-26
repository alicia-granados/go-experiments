package main

import "fmt"

func main() {

	x := 42

	if x == 40 {
		fmt.Println("Ell valor de x es 40")
	} else if x == 41 {
		fmt.Println("El valor de x es 41")
	} else if x == 42 {
		fmt.Println("El valor de x es 42")
	} else if x == 43 {
		fmt.Println("El valor de x es 43")
	} else {
		fmt.Println("El valo de x no es ninguno de estos números")
	}
}
