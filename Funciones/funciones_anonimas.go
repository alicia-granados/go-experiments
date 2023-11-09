package main

import (
	"fmt"
)

func saludar(name string, f func(string)) {
	f(name)
}

func duplicar(n int) int {
	return n * 2
}

func triplicar(n int) int {
	return n * 3
}

func aplicar(f func(int) int, n int) int {
	return f(n)
}

func main() {
	foo()
	// funcion anonima
	func() {
		fmt.Println("La función se ejecutó")
	}()

	func(x int) {
		fmt.Println("El significado de la vida es: ", x)
	}(42)

	saludo := func(name string) {
		fmt.Printf("Hola, %s\n", name)
	}
	saludo("juan")

	saludar("alex", saludo)

	result1 := aplicar(duplicar, 5)
	result2 := aplicar(triplicar, 4)

	fmt.Println(result1)
	fmt.Println(result2)

}

func foo() {
	fmt.Println("Foo se ejecutó")
}
