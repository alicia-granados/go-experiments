package main

import "fmt"

type canino struct {
	nombre string
	edad   int
}

func main() {
	fido := canino{
		nombre: "Fido",
		edad:   Edad(10),
	}

	fmt.Println(fido)
	fmt.Println(EdadDos(20))
}

// Edad Convierte en años perro
func Edad(n int) int {
	return n * 7
}

// EdadDos convierte años humanos en años perro
func EdadDos(n int) int {
	count := 0
	for i := 0; i < n; i++ {
		count += 7
	}
	return count
}

// go test
// go test -bench .
// coverage: go test -cover
// coverage mostrado en el navegador :
// go test -coverprofile c.out
// go tool cover -html=c.out
// godoc -http=:8080
