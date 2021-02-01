package main

import "fmt"

type person struct {
	nombre   string
	apellido string
	edad     int
}

type agenteSecreto struct {
	person
	lmp2 bool
}

func main() {

	as1 := agenteSecreto{

		person: person{
			nombre:   "Eduardo",
			apellido: "Tua",
			edad:     25,
		},
		lmp2: true,
	}

	p2 := person{
		nombre:   "Elias",
		apellido: "Pérez",
		edad:     15,
	}

	fmt.Println(as1)
	fmt.Println(p2)

	fmt.Println(as1.person.nombre, as1.apellido, as1.edad)
	fmt.Println(p2.nombre, p2.apellido, p2.edad)

}
