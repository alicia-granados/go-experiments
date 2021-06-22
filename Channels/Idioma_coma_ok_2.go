package main

import (
	"fmt"
)

func main() {

	// Select: la declaración select extrae  extrae el valor de cualquier canal que tenga un valor listo para ser extraido
	c := make(chan int)

	go func() {
		c <- 42
		close(c)
	}()

	v, ok := <-c
	fmt.Println(v, ok)

	v, ok = <-c
	fmt.Println(v, ok)

}
