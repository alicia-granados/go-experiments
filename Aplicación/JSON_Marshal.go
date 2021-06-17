package main

import (
	"encoding/json"
	"fmt"
)

type persona struct {
	Nombre   string
	Apellido string
	Edad     int
}

func main() {
	p1 := persona{
		Nombre:   "James",
		Apellido: "Bond",
		Edad:     22,
	}
	p2 := persona{
		Nombre:   "Mias",
		Apellido: "Moneypenny",
		Edad:     27,
	}
	personas := []persona{p1, p2}
	fmt.Println(personas)

	bs, err := json.Marshal(personas)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(bs))
}
