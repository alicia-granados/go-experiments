package main

import "fmt"

func PrintList(list ...interface{}) {
	for _, value := range list {
		fmt.Println(value)
	}
}

// es lo mismo que la funcion anterior solo que desde versiones actuales de go se puede usar any
func PrintList2(list ...any) {
	for _, value := range list {
		fmt.Println(value)
	}
}
func main() {
	PrintList("alex", "roel", "juan")
	PrintList(100, 200, 1, 132)
	PrintList2("alex", "roel", "juan")
	PrintList2(100, 200, 1, 132)
}
