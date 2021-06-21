package main

import (
	"fmt"
)

func main() {

	c := make(chan int) // canal general

	// enviar
	go enviar(c)

	//recibir
	recibir(c)

	fmt.Println("Finalizando")

}

func enviar(c chan<- int) {
	c <- 42
}

func recibir(c <-chan int) {
	fmt.Println(<-c)
}
