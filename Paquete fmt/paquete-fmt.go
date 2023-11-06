package main

import "fmt"

var a int
var b string = "Programa"
var c bool
var d bool = true

func main() {

	e := 42
	f := "dice hola mundo."
	g := `el doctor dice que comer vegetales es 
			saludable`
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
	fmt.Println("Mi", b, f)
	fmt.Println(g)

	/*Print F formatea dependiendo del identificador del formato*/
	fmt.Printf("El valor de la variable a es : %v\n", a)
	fmt.Printf("El valor de la variable b es : %v\n", b)
	fmt.Printf("El tipo de a es: %T\n", a)
	fmt.Printf("El tipo de b es: %T\n", b)

	s := fmt.Sprint("El ", b, " dice hola mundo")
	fmt.Println(s)
	fmt.Printf("%T", s)

	fmt.Print("Hola mundo sin salto de linea")

	name := "alex"
	age := 28

	fmt.Printf("hola me llamo %s y tengo %d años \n", name, age)

	greeting := fmt.Sprintf("hola me llamo %s y tengo %d años ", name, age)
	fmt.Println(greeting)

	var nombre, apellido string
	var edad int

	fmt.Println("Ingrese suy nombre:")
	fmt.Scanln(&nombre, &apellido)

	fmt.Println("Ingrese suy edad:")
	fmt.Scanln(&edad)

	fmt.Printf("Me llamo %s %s y tengo %d años", name, apellido, age)

	// %v se recomienda usar cuando no se sabe el tipo de dato que se va a ingresar
	// %T  tipo de dato

}
