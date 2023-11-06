package main

import (
	"fmt"
)

func main() {

	s1 := "Hola Mundo"
	s2 := `esta es una linea 
	partida.`
	fmt.Println(s1)
	fmt.Printf("El tipo de s1 es: %T\n", s1)

	fmt.Println(s2)
	fmt.Printf("El tipo de s2 es: %T\n", s2)

	fmt.Println("")

	bs := []byte(s1)
	fmt.Println(bs)
	fmt.Printf("El tipo de bs es: %T\n", bs)

	fmt.Println("")

	for i := 0; i < len(s1); i++ { //hexadecimal
		fmt.Printf("%#U ", s1[i])
	}

	fmt.Println("")

	for i, v := range s1 { //decimal
		fmt.Printf("En el índice %d el valor es %v\n", i, v)
	}

	fmt.Printf("Con el  verbo %q indico que se imprima el string %s", "%s", s1)

	fullName := "Alex Roel \t(alias \"roelcode\")\n"
	fmt.Println(fullName)

	s := "hola"
	fmt.Println(s[0]) // impirme el valor ASCII

}
