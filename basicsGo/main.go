package main

import (
	"basicsGo/data"
	"fmt"
)

var name = "Frontend master"

func init() {
	fmt.Println("a")
}
func init() {
	fmt.Println("B")
}

func calculateTax(price float32) (float32, float32) {
	return price * 0.09, price * 0.02
}
func calculateTaxWithName(price float32) (stateTax float32, cityTax float32) {
	stateTax = price * 0.09
	cityTax = price * 0.02
	return stateTax, cityTax
}

func main() {
	fmt.Println("Hello world")
	print("hello from go")

	printData()

	fmt.Println(data.MaxSpeed)

	fmt.Println(calculateTax(100))

	stateTax, cityTax := calculateTax(1000)
	fmt.Println(stateTax, cityTax)
	fmt.Println(calculateTaxWithName(300))
}
