package main

import (
	"fmt"
)

func main() {
	var i int

	for i <= 10 {
		fmt.Println(i)
		i++
	}

	for i2 := 1; i2 <= 10; i2++ {

		if i2 == 5 {
			// break
			continue
		}
		fmt.Println(i2)
	}
}
