package main

import "fmt"

const (
	_   = iota
	kb1 = 1 << (iota * 10)
	gb1 = 1 << (iota * 10)
	tb1 = 1 << (iota * 10)
)

func main() {
	a := 4
	fmt.Printf("%d\t\t%b\n", a, a)

	b := a << 1
	fmt.Printf("%d\t\t%b\n", b, b)

	kb := 1024
	gb := kb * 1024
	tb := gb * 1024

	fmt.Printf("%d\t\t%b\n", kb, kb)
	fmt.Printf("%d\t\t%b\n", gb, gb)
	fmt.Printf("%d\t%b\n", tb, tb)

	fmt.Printf("%d\t\t%b\n", kb1, kb1)
	fmt.Printf("%d\t\t%b\n", gb1, gb1)
	fmt.Printf("%d\t%b\n", tb1, tb1)

}
