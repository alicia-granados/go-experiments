package main

import (
	"fmt"
)

func main() {
	hello("juan")
	saludo := hello2("juan")
	fmt.Println((saludo))

	fmt.Println(cal3(2, 4))

	sum, mul := cal2(2, 4)
	fmt.Println("La suma es ", sum)
	fmt.Println("La multiplicacipn es ", mul)
}

func hello(name string) {
	fmt.Println("hola", name)
}

func hello2(name string) string {
	return "hola " + name
}

func cal(a, b int) int {
	return a + b
}

func cal2(a, b int) (int, int) {
	sum := a + b
	mul := a * b
	return sum, mul
}

func cal3(a, b int) (sum, mul int) {
	sum = a + b
	mul = a * b
	return
}
