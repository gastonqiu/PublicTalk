package main

import (
	"html/template"
	"os"
	"strings"
)

type Inventory struct {
	Item []Item
}

type Item struct {
	Name  string
	Count uint
}

func main() {
	inventory := Inventory{
		Item: []Item{
			{"black tea", 17},
			{"green tea", 10},
			{"cookie", 5},
			{"coffee", 0},
		},
	}

	text := `
		{{- range $item := .Item }}
			{{ if ne $item.Count 0 -}}
			Remaining {{ $item.Count }}{{ Quantifier $item.Name }} {{ $item.Name -}}
			{{ end -}}
		{{ end -}}
	`

	tmpl, err := template.New("test").Funcs(
		template.FuncMap{
			"Quantifier": func(s string) string {
				if strings.Contains(s, "tea") {
					return " bottles of"
				}
				if strings.Contains(s, "cookie") {
					return " bags of"
				}

				return ""
			},
		},
	).Parse(text)
	if err != nil {
		panic(err)
	}

	err = tmpl.Execute(os.Stdout, inventory)
	if err != nil {
		panic(err)
	}
}
