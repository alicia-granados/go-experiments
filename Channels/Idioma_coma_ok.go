package main

import (
	"fmt"
)

func main() {

	// Select: la declaración select extrae  extrae el valor de cualquier canal que tenga un valor listo para ser extraido
	par := make(chan int)
	impar := make(chan int)
	salir := make(chan bool)

	// enviar
	go enviar(par, impar, salir)

	//recibir
	recibir(par, impar, salir)

	fmt.Println("Finalizando")

}

// enviar
func enviar(par, impar chan<- int, salir chan<- bool) {
	for j := 0; j < 100; j++ {
		if j%2 == 0 {
			par <- j
		} else {
			impar <- j
		}
	}
	close(salir)
}

//recibir
func recibir(par, impar <-chan int, salir <-chan bool) {
	for {
		select {
		case v := <-par:
			fmt.Println("Desde el canal par: ", v)

		case v := <-impar:
			fmt.Println("Desde el canal impar: ", v)

		case i, ok := <-salir:
			if !ok {
				fmt.Println("Desde coma ok: ", i, ok)
				return
			} else {
				fmt.Println("Desde coma ok: ", i)
			}

		}
	}
}
