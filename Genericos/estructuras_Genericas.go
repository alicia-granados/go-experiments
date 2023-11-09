package main

import "fmt"

type Product[T uint | string] struct {
	Id    T
	Desc  string
	Price float32
}

func main() {
	product1 := Product[uint]{1, "Zapatos", 50}
	fmt.Println(product1)

	product2 := Product[string]{"AS-AS-DAS-DAS-D", "Zapatos", 50}
	fmt.Println(product2)

}
