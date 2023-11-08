package main

import (
	"fmt"
	"log"

	"github.com/alicia-granados/greetings"
)

func main() {
	log.SetPrefix("greetings")
	log.SetFlags(0)
	message, err := greetings.Hello("alex")

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(message)

	names := []string{"alex", "roel", "juan"}
	messages, err := greetings.Hellos(names)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(messages)

}
