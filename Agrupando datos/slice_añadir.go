package main

import "fmt"

func main() {

	x := []int{4, 5, 46, 47, 8}
	y := []int{5, 3, 2, 1}

	x = append(x, 7, 8, 96)
	fmt.Println(x)

	x = append(x, y...)
	fmt.Println(x)

}
