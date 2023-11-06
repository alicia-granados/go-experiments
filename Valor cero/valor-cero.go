package main

import (
	"fmt"
)

//VALOR PREDETEMINADO O VALOR CERO
/*Declarando la varible con el identificador y de tipo int*/
var y int

func main() {

	y = 21
	fmt.Println(y)

	var (
		defaultInt    int
		defaulUint    uint
		defaultFloat  float32
		defaultBool   bool
		defaultString string
	)

	fmt.Println(defaultInt, defaulUint, defaultFloat, defaultBool, defaultString)
}

/* VALORES CERO*
FALSE para booleanos
0 para integers
0.0 para floats
"" para strings
nil para poniters,funciones,interfaces,slices,channels, map*/
