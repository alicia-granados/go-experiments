package main

import "fmt"

func main() {

	if true {
		fmt.Println(001)
	}
	if false {
		fmt.Println(002)
	}
	if !true {
		fmt.Println(003)
	}
	if !false {
		fmt.Println(004)
	}

	if 42 == 42 {
		fmt.Println(005)
	}

	if x := 42; x == 42 {
		fmt.Println(006)
	}
}
