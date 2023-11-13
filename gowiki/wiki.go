package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

type Page struct {
	Title string
	Body  []byte
}

func (p *Page) save() error {
	filename := p.Title + ".txt"
	return os.WriteFile(filename, p.Body, 0600)
}

func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/view/"):]
	p, _ := loadPage(title)
	fmt.Fprintf(w, "<h1>%s</h1> <div>%s</div>", p.Title, p.Body)
}
func main() {
	/*
		p1 := &Page{Title: "TestPage", Body: []byte("Esta es una pagina de muestra")}

		p1.save()

		p2, _ := loadPage("TestPage")
		fmt.Println(string(p2.Body))
	*/

	//responder al cliente con un mensaje
	http.HandleFunc("/view/", viewHandler)

	//levantar servidor y regustrar el error
	log.Fatal(http.ListenAndServe(":8080", nil))
}
