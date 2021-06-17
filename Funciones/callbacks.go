package main

import (
	"fmt"
)

func main() {

	ii := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	s := suma(ii...)
	fmt.Println(s)

	s1 := sumaPares(suma, ii...)
	fmt.Println("El total de pares es de :", s1)

	s2 := sumaImpares(suma, ii...)
	fmt.Println("El total de impares es de :", s2)
}

func suma(xi ...int) int {
	total := 0
	for _, v := range xi {
		total += v
	}
	return total
}

func sumaPares(f func(xi ...int) int, vi ...int) int {
	var y []int
	for _, v := range vi {
		if v%2 == 0 {
			y = append(y, v)
		}
	}
	return f(y...)
}

func sumaImpares(f func(xi ...int) int, vi ...int) int {
	var y []int
	for _, v := range vi {
		if v%2 != 0 {
			y = append(y, v)
		}
	}
	return f(y...)
}
