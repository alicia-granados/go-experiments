package main

import (
	"fmt"
	"structMethods/data"
)

func main() {

	max := data.Instructor{Id: 3, LastName: "Firtman"} // new Instructor
	max.FirstName = "Maximiliano"
	fmt.Println(max.Print())
	/*
		kyle := data.NewInstructor("Kyle", "Simpson")
		fmt.Println(kyle.Print())
	*/
	goCourse := data.Course{Id: 2, Name: "Go fundamentals", Instructor: max}
	fmt.Printf("%v", goCourse)

	swiftWS := data.NewWorkshop("Swift with ios", max)
	fmt.Printf("%v", &swiftWS)

	var courses [2]data.Signable
	//courses[0] = goCourse
	courses[1] = swiftWS

	for _, course := range courses {
		fmt.Println(course)
	}

	var things []interface{}
	things = append(things, "34")
	things = append(things, 34)
	things = append(things, swiftWS)

	things[2].(data.Workshop).SignUp()
	print(things)
}
