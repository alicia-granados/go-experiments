package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	s := `clave123`
	bs, err := bcrypt.GenerateFromPassword([]byte(s), 4)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(s)
	fmt.Println(bs)
	fmt.Println(string(bs)) // hash

	claveLogin := `clave1234`
	err = bcrypt.CompareHashAndPassword(bs, []byte(claveLogin)) // -> Contraseña Invalida
	// err = bcrypt.CompareHashAndPassword(bs, []byte(`clave123`)) // ->  COntraseña Valida
	if err != nil {
		fmt.Println("Contraseña Invalida")
		return
	}
	fmt.Println("Contraseña Valida")
}
