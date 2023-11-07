package main

import "fmt"

type Persona struct {
	nombre string
	edad   int
	correo string
}

// metodo pertenece a una estrutura
func (p *Persona) diHola() {
	fmt.Println("hola mi nombre es ", p.nombre)
}

func main() {
	var x int = 10
	var p *int = &x

	fmt.Println(x)
	fmt.Println(p)

	editar(&x)
	fmt.Println(x)

	p2 := Persona{"alexa", 28, "alexa@gmail.com"}
	p2.diHola()
}

func editar(x *int) {
	*x = 20
}
