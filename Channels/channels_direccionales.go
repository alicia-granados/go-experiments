package main

import (
	"fmt"
)

func main() {

	// buffered channel ( canal con bufer)
	ca := make(chan int, 2)
	ca <- 42
	ca <- 43
	fmt.Println(<-ca) // 42
	fmt.Println(<-ca) // 43

	fmt.Println("-----")
	fmt.Printf("%T\n", ca)

}
