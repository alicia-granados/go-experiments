package main

import (
	"errors"
	"fmt"
	"strconv"
)

func divide(dividendo, divisor int) (int, error) {
	if divisor == 0 {
		// fmt.Errorf("No se puede diidir por cero")
		return 0, errors.New("NO se puede dividir por cero")
	}
	return dividendo / divisor, nil
}
func main() {
	str := "123"
	num, err := strconv.Atoi(str)

	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	fmt.Println("Numero", num)

	result, err := divide(10, 2)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Resultado", result)
}
