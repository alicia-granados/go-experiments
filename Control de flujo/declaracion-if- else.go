package main

import (
	"fmt"
	"time"
)

func main() {

	t := time.Now()

	hora := t.Hour()

	if hora < 12 {
		fmt.Println("¡Mañana!")
	} else if hora < 17 {
		fmt.Println("¡Tarde!")
	} else {
		fmt.Println("Noche")
	}

	if t2 := time.Now(); t2.Hour() < 12 {
		fmt.Println("¡Mañana!")
	} else if t2.Hour() < 17 {
		fmt.Println("¡Tarde!")
	} else {
		fmt.Println("Noche")
	}

}
