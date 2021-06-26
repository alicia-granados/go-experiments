package main

import (
	"fmt"
	"strings"
)

func main() {
	const s = "Nos preguntamos: ¿Quién soy yo para ser brillante..."

	xs := strings.Split(s, " ")

	for _, v := range xs {
		fmt.Println(v)
	}

	fmt.Printf("\n%s\n", Cat(xs))
	fmt.Printf("\n%s\n\n", Join(xs))

}

// Cat junta strings
func Cat(xs []string) string {
	str := xs[0]
	for _, v := range xs[1:] {
		str += " "
		str += v
	}
	return str
}

// Join  une strings
func Join(xs []string) string {
	return strings.Join(xs, " ")
}
