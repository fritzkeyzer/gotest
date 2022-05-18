package main

import (
	"embed"
	"html/template"
	"os"
)

//go:embed *.html
var files embed.FS

func main(){

	// Parsing a single template


	
	

	tmpl, _ := template.ParseFS(files, "hello.html")
	tmpl.Execute(os.Stdout, "World") // Hello World!
}