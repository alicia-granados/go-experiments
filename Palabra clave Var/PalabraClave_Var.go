package main

import "fmt"

var z = 41

func main() {

	x := 42 + 7
	y := "James Bond"
	fmt.Println(x)
	fmt.Println(y)

	x = 50
	fmt.Println(x)
	fmt.Println(z)

	numero()

}

func numero() {
	fmt.Println(z)
}
