package main

import (
	"basicsGo/data"
	"fmt"
)

var name = "Frontend master"

func main() {
	fmt.Println("Hello world")
	print("hello from go")

	printData()

	fmt.Println(data.MaxSpeed)
}
