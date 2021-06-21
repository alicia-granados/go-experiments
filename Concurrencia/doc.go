package main

import (
	"fmt"
)

func main() {
	ch := make(chan int)
	go func() {
		ch <- multiplica(5)
	}()
	fmt.Println(<-ch)
}

func multiplica(x int) int {
	return x * 5
}
