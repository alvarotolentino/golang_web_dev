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
	Functional bool
}

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	rust := language{Name: "Rust", Creator: "Mozilla", Functional: false}
	golang := language{Name: "Go", Creator: "Google", Functional: false}
	erlang := language{Name: "Erlang", Creator: "Ericson", Functional: true}
	elixir := language{Name: "Elixir", Creator: "José Valim", Functional: true}

	languages := []language{rust, golang, erlang, elixir}
	err := tpl.Execute(os.Stdout, languages)
	if err != nil {
		log.Fatalln(err)
	}
}
