package main

import (
	"fmt"
	"runtime"
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

	switch t := time.Now(); {
	case t.Hour() < 12:
		fmt.Println("¡Mañana!")
	case t.Hour() < 17:
		fmt.Println("¡Tarde!")
	default:
		fmt.Println("Noche")
	}
	os := runtime.GOOS
	switch os {
	case " windows":
		fmt.Println("GO run -> Windows")

	case " linux":
		fmt.Println("GO run -> 	Linux")

	case " darwin":
		fmt.Println("GO run -> 	macOS")
	default:
		fmt.Println("GO run -> otro OS")
	}
}
