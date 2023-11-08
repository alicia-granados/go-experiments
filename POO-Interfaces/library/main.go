package main

import (
	"fmt"
	"library/animal"

	"library/book"
)

func main() {

	var myBook = book.Book{
		Tittle: "Mody Dick",
		Author: "Herman Melville",
		Pages:  300,
	}
	myBook.PrintInfo()

	myBook2 := book.NewBook("Mody Dick", "Herman Melville", 300)
	myBook2.PrintInfo()

	myBook3 := book.NewBook2("Mody Dick", "Herman Melville", 300)
	myBook3.PrintInfo2()

	myBook3.SetTitle("moby dock( edicion especial)")
	fmt.Println(myBook3.GetTitle())

	myTextBook := book.NewTexBook("comunicacion", "jaime", 261, "sanrtillaana", "secundaria")
	myTextBook.PrintInfo3()

	//book.Print(myBook3)
	book.Print(myTextBook)

	miPerro := animal.Perro{Nombre: "max"}
	miGato := animal.Gato{Nombre: "Tom"}

	animal.HacerSonido(&miPerro)
	animal.HacerSonido(&miGato)

	animales := []animal.Animal{
		&animal.Perro{Nombre: "max"},
		&animal.Gato{Nombre: "Tom"},
		&animal.Perro{Nombre: "maxdf"},
		&animal.Gato{Nombre: "Tome"},
	}

	for _, animal := range animales {
		animal.Sonido()
	}
}
