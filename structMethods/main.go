package main

import (
	"fmt"
	"structMethods/data"
)

func main() {
	max := data.Instructor{Id: 3, LastName: "Firtman"} // new Instructor
	max.FirstName = "Maximiliano"
	fmt.Println(max.Print())
}
