package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

type language struct {
	Name    string
	Creator string
}

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	rust := language{Name: "Rust", Creator: "Mozilla"}

	err := tpl.Execute(os.Stdout, rust)
	if err != nil {
		log.Fatalln(err)
	}
}
