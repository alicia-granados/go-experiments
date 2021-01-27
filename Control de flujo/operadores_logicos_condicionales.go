package main

import "fmt"

func main() {

	fmt.Println(2 == 2 && 3 == 3)
	fmt.Println(true && false)
	fmt.Println(5 == 5 || 6 == 6)
	fmt.Println(true || false)
	fmt.Println(!true)

	fmt.Printf(" 2 == 2 && 3 == 3 \t %v\n", 2 == 2 && 3 == 3)
	fmt.Printf(" true && false \t \t %v\n", true && false)
	fmt.Printf(" 5 == 5 || 6 == 6 \t %v\n", 5 == 5 || 6 == 6)
	fmt.Printf(" true || false \t \t %v\n", true || false)
	fmt.Printf(" !true\t\t\t%v\n", !true)

}
