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

func birthday(pointerAge *int) {
	fmt.Println(pointerAge)  // memory address
	fmt.Println(*pointerAge) // value
	fmt.Printf("The pointer is %v and the value is %v\n", pointerAge, *pointerAge)

	if *pointerAge > 140 {
		panic("Too old to be true")
	}
	*pointerAge++
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

	defer fmt.Println("BYEEEEE!!!")
	defer fmt.Println("GOOD!!!")

	age := 22

	fmt.Println(age)
	birthday(&age)
	fmt.Println(age)
	fmt.Println(&age) // memory address
}
