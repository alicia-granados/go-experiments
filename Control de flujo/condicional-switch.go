package main

import "fmt"

func main() {

	switch {
	case 2 == 4, 4 == 8, 9 == 9:
		fmt.Println("No debería imprimir,ahora sí")
	case 3 == 3:
		fmt.Println("Debería imprimir")
		fallthrough
	case 4 == 5:
		fmt.Println("No debería imprimir 2")
	default:
		fmt.Println("Impriiendo desde el default")
	}

	fmt.Println("")

	switch "Manzana" {
	case "Pera", "Limón":
		fmt.Println("Frutas verde(s.")
	case "Manzana", "Ciruela", "Fresas":
		fmt.Println("Frutas rojas")
		fallthrough
	default:
		fmt.Println("Imprimiendo desde el default")
	}

	fmt.Println("")
	switch {
	case false:
		fmt.Println("Falso")
	case true:
		fmt.Println("Verdadero")
	}

}
