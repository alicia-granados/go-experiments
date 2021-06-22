package main

import (
	"context"
	"fmt"
)

func main() {

	ctx := context.Background()

	fmt.Println("context:\t\t", ctx)
	fmt.Println("Error de context:\t", ctx.Err())
	fmt.Printf("Tipo de context:\t%T\n", ctx)
	fmt.Println("-----------")

	ctx, cancel := context.WithCancel(ctx)

	fmt.Println("context:\t\t", ctx)
	fmt.Println("Error de context:\t", ctx.Err())
	fmt.Printf("Tipo de context:\t%T\n", ctx)
	fmt.Println("Cancel:\t\t\t", cancel)
	fmt.Printf("Tipo de cancel:\t\t%T\n", cancel)
	fmt.Println("-----------")

	cancel()

	fmt.Println("context:\t\t", ctx)
	fmt.Println("Error de context:\t", ctx.Err())
	fmt.Printf("Tipo de context:\t%T\n", ctx)
	fmt.Println("Cancel:\t\t\t", cancel)
	fmt.Printf("Tipo de cancel:\t\t%T\n", cancel)
	fmt.Println("-----------")

}
