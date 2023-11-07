package main

import (
	"fmt"
	"os"
)

func main() {
	defer fmt.Println(3)
	defer fmt.Println(1)
	defer fmt.Println(2)

	file, err := os.Create("Hola.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	_, err = file.Write([]byte("Hola alex roel"))

	if err != nil {
		fmt.Println(err)

		return
	}

}
