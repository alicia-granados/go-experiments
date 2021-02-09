package main

import "fmt"

type persona struct {
	nombre   string
	apellido string
}

type agenteSecreto struct {
	persona
	lpm bool
}

func (a agenteSecreto) presentar() {
	fmt.Println("Hola,  soy ", a.nombre, a.apellido, "el agente secreto se presenta")
}

func (p persona) presentar() {
	fmt.Println("Hola,  soy ", p.nombre, p.apellido, "la persona se presenta")
}

type humano interface {
	presentar()
}

func met2(h humano) {

	switch h.(type) {
	case persona:
		fmt.Println("Fui pasando a la función met2 (persona) ", h.(persona).nombre)
	case agenteSecreto:
		fmt.Println("Fui pasando a la función met2 (agenteSecreto) ", h.(agenteSecreto).nombre)
	}

	fmt.Println("Fui pasando a la función met2", h)
}

type manzana int

func main() {

	as1 := agenteSecreto{
		persona: persona{
			nombre:   "Eduard",
			apellido: "TUa",
		},
		lpm: true,
	}

	as2 := agenteSecreto{
		persona: persona{
			nombre:   "Condor",
			apellido: "Pérez",
		},
		lpm: true,
	}

	p := persona{
		nombre:   "Carito",
		apellido: "Guz",
	}

	fmt.Println(as1)
	as1.presentar()
	as2.presentar()
	met2(as1)
	met2(as2)
	met2(p)
	//conversión
	var x manzana = 42
	fmt.Println(x)
	fmt.Printf("%T\n", x)

	var y int
	y = int(x)
	fmt.Println(y)
	fmt.Printf("%T\n", y)

}
