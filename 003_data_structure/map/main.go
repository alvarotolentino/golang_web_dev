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
	languages := map[string]string{
		"Golang": "Google",
		"Rust":   "Mozilla",
		"Erlang": "Ericson",
		"Elixir": "José Valim"}

	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", languages)
	if err != nil {
		log.Fatalln(err)
	}
}
