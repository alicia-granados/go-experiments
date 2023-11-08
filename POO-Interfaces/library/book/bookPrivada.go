package book

import (
	"fmt"
)

type Printable interface {
	PrintInfo3()
}

func Print(p Printable) {
	p.PrintInfo3()
}

type BookPrivada struct {
	tittle string
	author string
	pages  int
}

func NewBook2(tittle string, author string, pages int) *BookPrivada {
	return &BookPrivada{
		tittle: tittle,
		author: author,
		pages:  pages,
	}
}

func (b *BookPrivada) SetTitle(title string) {
	b.tittle = title
}

func (b *BookPrivada) GetTitle() string {
	return b.tittle
}

// metodo
func (b *BookPrivada) PrintInfo2() {
	fmt.Printf("Tttle: %s\nauthor: %s\npages: %d\n", b.tittle, b.author, b.pages)
}

type TexBook struct {
	BookPrivada
	editorial string
	level     string
}

func NewTexBook(title, author string, pages int, editorial, level string) *TexBook {
	return &TexBook{
		BookPrivada: BookPrivada{title, author, pages},
		editorial:   editorial,
		level:       level,
	}
}

func (b *TexBook) PrintInfo3() {
	fmt.Printf("Tttle: %s\nauthor: %s\npages: %d\n editorial: %s\n nivel: %s", b.tittle, b.author, b.pages, b.editorial, b.level)
}
