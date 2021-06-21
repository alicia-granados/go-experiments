package main

import (
	"fmt"
)

func main() {

	//Range _ deja de leer desde un cnal cuando el canal es cerrado
	c := make(chan int) // canal general

	// enviar
	go func() {
		for i := 0; i < 5; i++ {
			c <- i
		}
		close(c)
	}()
	//recibir
	for v := range c {
		fmt.Println(v)
	}

	fmt.Println("Finalizando")

}
