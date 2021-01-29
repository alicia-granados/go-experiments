package main

import "fmt"

func main() {

	x := make([]int, 10, 12) // si su capacidad  se excede al ingresar datos se actualizara la capacidad al doble

	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

	x[0] = 1
	x[5] = 100

	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

	x = append(x, 200)
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

	x = append(x, 300)
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

	x = append(x, 400)
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

}
