package main

import "fmt"

func main() {
	// declaracion de variables
	var firstName, lastName string
	var age int

	firstName = "Alex"
	lastName = "Roel"
	age = 27

	fmt.Println(firstName, lastName, age)

	var (
		firstName2 = "alex"
		lastName2  = "roel"
		age2       = 27
	)
	fmt.Println(firstName2, lastName2, age2)

	var firstName3, lastName3, age3 = "alex", "roel", 27
	fmt.Println(firstName3, lastName3, age3)

	//Sin utilizar la palabra reservada dentro de main
	firstName4, lastName4, age4 := "alex", "roel", 27
	fmt.Println(firstName4, lastName4, age4)
}
