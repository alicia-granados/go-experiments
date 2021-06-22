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

	ctx, _ = context.WithCancel(ctx)

	fmt.Println("context:\t\t", ctx)
	fmt.Println("Error de context:\t", ctx.Err())
	fmt.Printf("Tipo de context:\t%T\n", ctx)
	fmt.Println("-----------")

}
