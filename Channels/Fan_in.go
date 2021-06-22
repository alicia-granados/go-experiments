package main

import (
	"fmt"
	"sync"
)

func main() {

	par := make(chan int)
	impar := make(chan int)
	fanin := make(chan int)

	// enviar
	go enviar(par, impar)

	//recibir
	go recibir(par, impar, fanin)

	for v := range fanin {
		fmt.Println(v)
	}

	fmt.Println("Finalizando")

}

// enviar
func enviar(par, impar chan<- int) {
	for j := 0; j < 100; j++ {
		if j%2 == 0 {
			par <- j
		} else {
			impar <- j
		}
	}
	close(par)
	close(impar)
}

//recibir
func recibir(par, impar <-chan int, fanin chan<- int) {
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		for v := range par {
			fanin <- v
		}
		wg.Done()
	}()

	go func() {
		for v := range impar {
			fanin <- v
		}
		wg.Done()
	}()
	wg.Wait()
	close(fanin)
}
