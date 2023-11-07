package main

import "fmt"

func divide(dividendo, divisor int) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()
	validateZero(divisor)
	fmt.Println(dividendo / divisor)
}

func validateZero(divisor int) {
	if divisor == 0 {
		panic("No PUEDES DIVIDIR POR CERO")
	}
}

func main() {
	divide(100, 10)
	divide(100, 0)
	divide(100, 1)
}
