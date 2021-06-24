package main

import "fmt"

func main() {
	fmt.Println("La suma de 2 + 4 es : ", miSuma(2, 4))
	fmt.Println("La suma de 1 + 5 es : ", miSuma(1, 5))
	fmt.Println("La suma de 12 + 4 es : ", miSuma(12, 4))
}

func miSuma(xi ...int) int {
	var sum int
	for _, v := range xi {
		sum += v
	}
	return sum
}
