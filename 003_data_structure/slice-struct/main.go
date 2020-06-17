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
	golang := language{Name: "Go", Creator: "Google"}
	erlang := language{Name: "Erlang", Creator: "Ericson"}
	elixir := language{Name: "Elixir", Creator: "José Valim"}

	languages := []language{rust, golang, erlang, elixir}
	err := tpl.Execute(os.Stdout, languages)
	if err != nil {
		log.Fatalln(err)
	}
}
