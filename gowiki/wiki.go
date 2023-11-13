package main

import (
	"html/template"
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
	renderTemplates(w, "view", p)

}

func editHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/edit/"):]
	p, _ := loadPage(title)
	renderTemplates(w, "edit", p)
}

func renderTemplates(w http.ResponseWriter, tmpl string, p *Page) {
	//cargar plantilla html
	t, _ := template.ParseFiles(tmpl + ".html")
	//ejecutar
	t.Execute(w, p)
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
	http.HandleFunc("/edit/", editHandler)
	//levantar servidor y regustrar el error
	log.Fatal(http.ListenAndServe(":8082", nil))
}
