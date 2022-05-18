package main

import (
	"bytes"
	"github.com/alecthomas/chroma/quick"
	"html/template"
	"log"
	"net/http"
)

type CodeSnip struct {
	Code   string
	Status bool
}

type Page struct {
	Title string
	Body  template.HTML
}

func main() {

	someSourceCode := `func main(){
	log.Println("hello world: %v", []string{"hi", "mom"})
}`

	buf := new(bytes.Buffer)

	err := quick.Highlight(buf, someSourceCode, "go", "html", "monokai")
	if err != nil {
		log.Fatalln(err)
	}

	// run template and serve on localhost
	tmpl := template.Must(template.ParseFiles("layout.html"))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data := Page{
			Title: "My chroma test",
			Body:  template.HTML(buf.String()),
		}
		tmpl.Execute(w, data)
	})
	http.ListenAndServe(":1337", nil)
}
