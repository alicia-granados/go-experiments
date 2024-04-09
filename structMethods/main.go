package main

import (
	"fmt"
	"structMethods/data"
)

func main() {
	max := data.Instructor{Id: 3, LastName: "Firtman"} // new Instructor
	max.FirstName = "Maximiliano"
	fmt.Println(max.Print())

	kyle := data.NewInstructor("Kyle", "Simpson")
	fmt.Println(kyle.Print())

	goCourse := data.Course{Id: 2, Name: "Go fundamentals", Instructor: max}
	fmt.Printf("%v", goCourse)
}
