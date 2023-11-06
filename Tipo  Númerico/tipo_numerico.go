package main

import (
	"fmt"
	"math"
	"runtime"
)

var w int
var z uint8

func main() {

	x := 42
	y := 42.532

	fmt.Println(x)
	fmt.Println(y)

	fmt.Printf("%T\n", x)
	fmt.Printf("%T\n", y)

	w = 4
	z = uint8(w) // valores positivos

	fmt.Println(w)
	fmt.Println(z)

	fmt.Printf("%T\n", w)
	fmt.Printf("%T\n", z)

	fmt.Println(runtime.GOOS)
	fmt.Println(runtime.GOARCH)

	var integer int8 = 127
	fmt.Println((integer))

	var float float32 // numeros pequeños
	fmt.Println(float)

	fmt.Println(math.MaxFloat32)

	//  valor ascii
	var a byte = 'a'
	fmt.Println(a)

}
