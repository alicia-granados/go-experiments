package main

import (
	"fmt"
	"math"
)

func main() {

	var lado1, lado2 float64
	const precision = 2

	//entrada y escaneada de datos
	fmt.Println("Ingrese lado 1 :")
	fmt.Scanln(&lado1)
	fmt.Println("Ingrese lado 2")
	fmt.Scanln(&lado2)

	aCuadrada := math.Pow(lado1, 2)
	fmt.Printf("A^2 = %.2f\n", aCuadrada)
	bCuadrada := math.Pow(lado2, 2)
	fmt.Printf("B^2 = %.2f\n", bCuadrada)
	// c^2 = a^2+ b^2

	hipotenusaCuadrada := aCuadrada + bCuadrada

	c := math.Sqrt(hipotenusaCuadrada)
	fmt.Printf("Hipotenusa %.2f \n", c)

	area := (lado1 * lado2) / 2
	perimetro := lado1 + lado2 + c

	fmt.Printf("Area: %.*f\n", precision, area)
	fmt.Printf("Perimetro: %.*f\n", precision, perimetro)

}
