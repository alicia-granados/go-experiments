package book

import "fmt"

type Book struct {
	Tittle string
	Author string
	Pages  int
}

func NewBook(tittle string, author string, pages int) *Book {
	return &Book{
		Tittle: tittle,
		Author: author,
		Pages:  pages,
	}
}

// metodo
func (b *Book) PrintInfo() {
	fmt.Printf("Tttle: %s\nAuthor: %s\nPages: %d\n", b.Tittle, b.Author, b.Pages)
}
