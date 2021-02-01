package main

import "fmt"

type persona struct {
	first string
	last  string
	age   int
}

func main() {

	p1 := persona{
		first: "Eduardo",
		last:  "Tua",
		age:   28,
	}

	p2 := struct {
		first string
		last  string
		age   int
	}{
		first: "Elias",
		last:  "Pérez",
		age:   15,
	}

	fmt.Println(p1)
	fmt.Println(p2)
}
