package main

import (
	"fmt"
)

func main() {
	x := 0
	fmt.Println("x antes :", x)
	fmt.Println("x antes :", &x)
	foo(&x)
	fmt.Println("x después: ", x)
	fmt.Println("x después: ", &x)
}

func foo(y *int) {
	fmt.Println("y antes:", y)
	fmt.Println("y antes:", *y)
	*y = 42
	fmt.Println("y después:", y)
	fmt.Println("y después:", *y)
}
