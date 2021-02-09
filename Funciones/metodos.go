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

func (a agenteSecreto) presentarse() {
	fmt.Println("Hola,  soy ", a.nombre, a.apellido)
}
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

	fmt.Println(as1)

	as1.presentarse()
	as2.presentarse()
}
