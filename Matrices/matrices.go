package main

import "fmt"

func main() {
	var matriz [5]int
	fmt.Println(matriz)

	var matriz2 = [5]int{10, 20, 30, 40, 50}
	matriz2[0] = 100
	fmt.Println(matriz2)

	var matriz3 = [...]int{10, 20, 30, 40} // ... significa que no se sabe con cuantos datos tendra la matriz
	fmt.Println(matriz3[1])

	for i := 0; i < len(matriz3); i++ {
		fmt.Println(matriz3[i])
	}

	for index, value := range matriz3 {
		fmt.Println("indice: %d, valor: %d\n, index, value", index, value)
	}

	var matrizBidimensional = [3][3]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	fmt.Println(matrizBidimensional)
}
