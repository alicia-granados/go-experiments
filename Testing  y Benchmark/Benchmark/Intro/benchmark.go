package main

import "fmt"

func main() {
	fmt.Println(Saludar("Eduardo"))
}

// saludar nos deja saludar a una persona
func Saludar(s string) string {
	return fmt.Sprint("Bienvenido querido ", s)
}

// Benchmarking: parte del paquete testing nos permite medir la velocidad de nuestro código
// go test -bench . muestra el tiempo de operación
// go test -bench Nombredelafuncion
