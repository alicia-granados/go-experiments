package main

import (
	"fmt"
)

func main() {

	// Select: la declaración select extrae  extrae el valor de cualquier canal que tenga un valor listo para ser extraido
	par := make(chan int)
	impar := make(chan int)
	salir := make(chan int)

	// enviar
	go enviar(par, impar, salir)

	//recibir
	recibir(par, impar, salir)

	fmt.Println("Finalizando")

}

// enviar
func enviar(p, i, s chan<- int) {
	for j := 0; j < 100; j++ {
		if j%2 == 0 {
			p <- j
		} else {
			i <- j
		}
	}
	//close(p)
	//close(i)
	s <- 0
}

//recibir
func recibir(p, i, s <-chan int) {
	for {
		select {
		case v := <-p:
			fmt.Println("Desde el canal par: ", v)

		case v := <-i:
			fmt.Println("Desde el canal impar: ", v)

		case v := <-s:
			fmt.Println("Desde el canal salir: ", v)
			return
		}
	}
}
