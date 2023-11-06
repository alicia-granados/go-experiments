package main

import (
	"fmt"
	"strconv"
)

var a int

type dinero int

var b dinero

func main() {
	a = 42
	fmt.Println(a)
	fmt.Printf("%T\n", a)

	b = 1000
	fmt.Println(b)
	fmt.Printf("%T\n", b)

	a = int(b)
	fmt.Println(a)
	fmt.Printf("%T\n", a)

	var integer16 int16 = 50
	var integer32 int32 = 100

	fmt.Println(integer16 + int16(integer32))

	//conversion de string a entero
	s := "100"
	i, _ := strconv.Atoi(s)
	fmt.Println(i + 1)

	//conversion de int a string
	n := 42
	s = strconv.Itoa(n)
	fmt.Println(s)

}
