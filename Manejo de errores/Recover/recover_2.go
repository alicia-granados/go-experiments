package main

import (
	"fmt"
)

func main() {
	var x int
	x++
	fmt.Println(x)
	v := c()
	fmt.Println(v)
}

func c() (i int) {
	defer func() { i++ }()
	return 1
}
