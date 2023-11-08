package book

import (
	"fmt"
)

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
