package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {

	fmt.Println("OS\t\t", runtime.GOOS)                 // Sistema Operativo
	fmt.Println("AECH\t\t", runtime.GOARCH)             // Arquitectura
	fmt.Println("CPUs\t\t", runtime.NumCPU())           // Número de procesadores que se utilizan
	fmt.Println("GOroutines\t", runtime.NumGoroutine()) // Número de gorutinas

	wg.Add(1) // Número de rutinas que habra

	go foo()
	bar()
	wg.Wait() // Agrega metodo  a main para que espere a las go rutines terminen

}

func foo() {
	for i := 0; i < 10; i++ {
		fmt.Println("foo:", i)
	}
	wg.Done() //A  cada go rutina se le agrega ese método para que informe a main que ya realizo su trabajo  y se le resta a ADD
}

func bar() {
	for i := 0; i < 10; i++ {
		fmt.Println("bar:", i)
	}
}
