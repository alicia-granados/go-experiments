package main

import "fmt"

// declaracion  de constantes
const Pi float32 = 3.14
const Pi2 = 3.14

const (
	x = 100
	y = 0b1010 // Binario
	z = 0o12   // Ocatal
	w = 0xFF   //  Hexadecimal
)

const (
	Domingo = iota // iota inicia en cero
	Lunes
	Martes
	Miercoles
	Jueves
	Sabado
	Viernes
)

func main() {
	fmt.Println(Pi)
	fmt.Println(Pi2)
	fmt.Println(x, y, z, w)
	fmt.Println((Viernes))
}
