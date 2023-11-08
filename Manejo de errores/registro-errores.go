package main

import (
	"log"
	"os"
)

func main() {

	//log.Panic("este es otro mensaje de registrp")
	log.SetPrefix("main():")
	//log.Println("este es otro mensaje de registrp")
	//log.Fatal("Oye soy un registor de errores")

	file, err := os.OpenFile("info.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 044)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	log.SetOutput(file)
	log.Print("oye soy un log")
}
