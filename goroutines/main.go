package main

import (
	"fmt"
	"time"
)

func a() {

}

func printMessage(text string) {
	go a()
	for i := 0; i < 5; i++ {
		fmt.Println(text)
		time.Sleep(800 * time.Microsecond)
	}
}
func main() {
	go printMessage("go is great")
	go printMessage("Frontend Masters Rocks!")
	go printMessage("we miss classes!")
	time.Sleep(time.Minute)
}
