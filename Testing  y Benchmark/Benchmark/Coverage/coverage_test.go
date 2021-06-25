package main

import (
	"fmt"
	"testing"
)

func TestSaludar(t *testing.T) {
	s := Saludar("Eduardo")

	if s != "Bienvenido querido Eduardo" {
		t.Error("Expected", "Bienvenido querido Eduardo", "Got", s)
	}
}

func ExampleSaludar() {
	fmt.Println(Saludar("Eduardo"))
	//output:
	//Bienvenido querido Eduardo
}

func BenchmarkSaludar(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Saludar("Eduardo")
	}
}
