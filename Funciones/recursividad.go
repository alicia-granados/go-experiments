package main

import (
	"fmt"
)

func main() {
	a := 4
	fmt.Println(factorial(a))
	fmt.Println(ciclo_factorial(a))
}

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

func ciclo_factorial(n int) int {
	total := 1
	for ; n > 0; n-- {
		total *= n
	}
	return total
}
