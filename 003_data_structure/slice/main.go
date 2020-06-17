package main

import (
	"html/template"
	"log"
	"os"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	languages := []string{"Golang", "Rust", "Erlang", "Elixir"}
	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", languages)
	if err != nil {
		log.Fatalln(err)
	}
}
