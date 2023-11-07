package main

import "fmt"

func main() {
	// map[tipo]tipo de dato que va a ser su valor
	colors := map[string]string{
		"rojo":  "#FF0000",
		"verde": "#00FF00",
		"azul":  "#0000FF",
	}

	fmt.Println(colors["rojo"])

	colors["negro"] = "000000"
	fmt.Println(colors)

	valor := colors["rojo"]
	fmt.Println(valor)

	valor2, verificacion := colors["blanco"]
	fmt.Println(valor2, verificacion)

	if verificacion {
		fmt.Println("Si existe la clave")
	} else {
		fmt.Println("NO existe esta clave")
	}

	fmt.Println(valor2, verificacion)

	if valor3, verificacion := colors["azul"]; verificacion {
		fmt.Println(valor3)
	} else {
		fmt.Println("NO existe esta clave")
	}

	delete(colors, "azul")
	fmt.Println(colors)

	for clave, valor4 := range colors {
		fmt.Printf("Clave: %s, valor: %s \n ", clave, valor4)
	}
}
