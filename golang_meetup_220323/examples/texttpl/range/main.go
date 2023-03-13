package main

import (
	"html/template"
	"os"
)

type Inventory struct {
	Items []Item
}

type Item struct {
	Name  string
	Count uint
}

func main() {
	inventory := Inventory{
		Items: []Item{
			{"black tea", 17},
			{"green tea", 10},
			{"cookie", 5},
			{"coffee", 0},
		},
	}

	text := `
		{{- range $item := .Items }}
			Remaining {{ $item.Count }} {{ $item.Name -}}
		{{ end -}}
	`

	tmpl, err := template.New("test").Parse(text)
	if err != nil {
		panic(err)
	}

	err = tmpl.Execute(os.Stdout, inventory)
	if err != nil {
		panic(err)
	}
}
