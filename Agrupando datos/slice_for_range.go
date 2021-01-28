package main

import "fmt"

func main() {

	x := []int{2, 3, 4, 42}
	fmt.Println(x)

	fmt.Println(cap(x)) //capacidad

	fmt.Println(x[0])
	fmt.Println(x[1])
	fmt.Println(x[2])
	fmt.Println(x[3])

	for i, v := range x {
		fmt.Println(i, v)
	}

}
