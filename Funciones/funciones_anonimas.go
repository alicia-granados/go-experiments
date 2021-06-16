package main

import (
	"fmt"
)

func main() {
	foo()

	func() {
		fmt.Println("La función se ejecutó")
	}()

	func(x int) {
		fmt.Println("El significado de la vida es: ", x)
	}(42)

}

func foo() {
	fmt.Println("Foo se ejecutó")
}
