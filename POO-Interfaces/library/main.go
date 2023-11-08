package main

import "library/book"

func main() {
	var myBook = book.Book{
		Tittle: "Mody Dick",
		Author: "Herman Melville",
		Pages:  300,
	}
	myBook.PrintInfo()

	myBook2 := book.NewBook("Mody Dick", "Herman Melville", 300)
	myBook2.PrintInfo()

}
