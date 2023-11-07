package main

import "fmt"

func main() {
	diaSemana := []string{"Domingo", "lunes", "martes", "miercoles", "jueves", "viernes", "sabado"}
	fmt.Println(diaSemana)

	diaRebanada := diaSemana[0:5]
	fmt.Println(diaRebanada)

	diaRebanada = append(diaRebanada, "viernes", "sabado", "otrodia")
	fmt.Println(diaRebanada)

	// elimina un dia y se agrega otro slice
	diaRebanada = append(diaRebanada[:2], diaRebanada[3:]...)
	fmt.Println(diaRebanada)

	fmt.Println(len(diaRebanada))
	fmt.Println(cap(diaRebanada))

	// crear slice mediamte make
	nombre := make([]string, 5, 10)
	nombre[0] = "alex"
	fmt.Println(nombre)

	//copiar elementos de una rebanada a otra
	rebanada1 := []int{1, 2, 3, 4, 5}
	rebanada2 := make([]int, 5)

	//copy(rebanada1, rebanada2) // no copia datos
	// copia de rebanada 1 a rebanada 2
	copy(rebanada2, rebanada1)
	fmt.Println(rebanada1)
	fmt.Println(rebanada2)
}
