package main

import "fmt"

// funcion variadica con un tipo de dato
func suma(nums ...int) int {

	var total int
	for _, num := range nums {
		total += num
	}
	//fmt.Printf("%T- %v\n", nums, nums)
	return total
}

// funcion variadica con un tipo de dato y n variables
func suma2(name string, nums ...int) int {

	var total int
	for _, num := range nums {
		total += num
	}
	fmt.Printf("Hola %s, la suma es %d\n", name, total)
	return total
}

// funcion variadica para n cantidad de datos y n tipo de datos
func imprimirDatos(datos ...interface{}) {
	for _, dato := range datos {
		fmt.Printf("%T -%v\n", dato, dato)
	}
}

func main() {
	fmt.Println(suma(12, 45, 30, 15))
	fmt.Println(suma2("alex", 1, 45, 302, 115, 1, 2, 3))
	imprimirDatos("hola", 20, 12.3, true, "s")
}
