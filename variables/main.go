package main

import "fmt"

// global variable
var url = "https://frontedmaster.coms"

func main() {
	// function-scoped variables
	var message string
	message = "Hello from go"
	print(message)

	var price float32 = 34.4
	print(message, price)

	price2 := 34.4
	print(price2)

	const maxSpeed byte = 60
	print(maxSpeed)

	fmt.Println(url)
	var isREady bool
	fmt.Println(isREady)

}
