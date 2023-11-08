package main

import (
	"fmt"
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
}
