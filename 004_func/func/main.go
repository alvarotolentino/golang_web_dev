package main

import (
	"log"
	"os"
	"strings"
	"text/template"
)

var tpl *template.Template

type language struct {
	Name    string
	Creator string
}

type car struct {
	Manufactured string
	Model        string
	Doors        int
}

var fm = template.FuncMap{
	"uc": strings.ToUpper,
	"ft": firstThree,
}

func init() {
	tpl = template.Must(template.New("").Funcs(fm).ParseFiles("tpl.gohtml"))
	// tpl = template.Must(template.ParseFiles("tpl.gohtml"))
	// tpl = tpl.Funcs(fm)
}

func firstThree(s string) string {
	s = strings.TrimSpace(s)
	s = s[:3]
	return s
}

func main() {
	a := language{
		Name:    "Erlang",
		Creator: "Ericson",
	}
	b := language{
		Name:    "Rust",
		Creator: "Mozilla",
	}

	x := car{
		Manufactured: "Ford",
		Model:        "F150",
		Doors:        2,
	}

	y := car{
		Manufactured: "Toyota",
		Model:        "Corolla",
		Doors:        4,
	}

	languages := []language{a, b}
	cars := []car{x, y}

	data := struct {
		Languages      []language
		Transportation []car
	}{
		Languages:      languages,
		Transportation: cars,
	}
	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", data)
	if err != nil {
		log.Fatalln(err)
	}
}
