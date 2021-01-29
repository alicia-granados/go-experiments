package main

import "fmt"

func main() {

	et := []string{"Eduardo", "López", "Crossfit", "Baseball", "Montañismo"}
	fmt.Println(et)

	j1 := []string{"Jacinto", "Tua", "Correr", "Nadar", "Bailar"}
	fmt.Println(j1)

	s1 := []string{"Norma", "Garcia", "Correr", "Nadar", "Bailar"}
	fmt.Println(j1)

	vp := [][]string{et, j1, s1}
	fmt.Println(vp)
}
