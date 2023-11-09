package main

import (
	"fmt"

	"golang.org/x/exp/constraints" // resctricciones que ya estan definidas
)

// restriccion de aproximacion  signo  -> ~ para que no pongan otros tipo de datos
type integer int

type Numbers interface {
	~int | ~float64 | ~float32 | ~uint
}

// funcion variadica con un tipo de dato y resticcion int o float
func suma[T ~int | ~float64](nums ...T) T {

	var total T
	for _, num := range nums {
		total += num
	}
	return total
}

func suma2[T Numbers](nums ...T) T {

	var total T
	for _, num := range nums {
		total += num
	}
	return total
}

func suma3[T constraints.Integer | constraints.Float](nums ...T) T {

	var total T
	for _, num := range nums {
		total += num
	}
	return total
}
func main() {
	fmt.Println(suma(4, 5, 6, 8))
	// ya no es necesario poner esto
	//fmt.Println(suma[int](4, 5, 6, 8))
	fmt.Println(suma(4.1, 5.5, 6.0, 8.8))

	var num1 integer = 100
	var num2 integer = 200
	fmt.Println(suma(num1, num2))
	fmt.Println(suma2(num1, num2))
	fmt.Println(suma3[float32](6.5, 4.3, 50.0))
}
