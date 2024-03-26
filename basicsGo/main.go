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
func main() {
	fmt.Println("Hello world")
	print("hello from go")

	printData()

	fmt.Println(data.MaxSpeed)
}
