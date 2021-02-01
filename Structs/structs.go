package main

import "fmt"

/*creamos valores de ciertos tipos que se almacenan en variables
las variables tienen identificadores*/

var x int

type persona struct {
	nombre   string
	apellido string
}

type foo int //Alias de tipos es mala práctica
var y foo

func main() {

	y = 42
	p1 := persona{
		nombre:   "Eduar",
		apellido: "Tua",
	}

	fmt.Println(p1)
	fmt.Printf("%T", y)

}
