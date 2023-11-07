package main

import "fmt"

type Persona struct {
	nombre string
	edad   int
	correo string
}

func main() {

	var p Persona
	p.nombre = " alex"
	p.edad = 28
	p.correo = "alex@gamil.com"

	fmt.Println(p)

	p2 := Persona{"alexa", 28, "alexa@gmail.com"}
	p.edad = 20
	fmt.Println(p2.nombre)

}
