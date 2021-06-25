package main

import "fmt"

func main() {
	fmt.Println(Saludar("Eduardo"))
}

// saludar nos deja saludar a una persona
func Saludar(s string) string {
	return fmt.Sprint("Bienvenido querido ", s)
}

//Muestra la covertura de los test
//go test -cover
//go test -coverprofile c.out
//mostrando en el navegador
// go tool cover -html=c.out
