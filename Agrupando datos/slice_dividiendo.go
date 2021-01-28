package main

import "fmt"

func main() {

	x := []int{4, 5, 7, 8, 42}
	fmt.Println(x)

	fmt.Println(x[1:4])
	fmt.Println(x[:])

	for i := 0; i <= 4; i++ {
		fmt.Println(i, x[i])
	}

}
