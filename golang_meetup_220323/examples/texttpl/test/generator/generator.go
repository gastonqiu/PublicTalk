package generator

import (
	"bytes"
	"html/template"
	"os"
	"strings"
)

type Inventory struct {
	Items []Item
}

type Item struct {
	Name  string
	Count uint
}

func Generate(inventory *Inventory) (string, error) {
	b, err := os.ReadFile("../_templates/text.tpl")
	if err != nil {
		return "", err
	}

	tmpl, err := template.New("test").Funcs(
		template.FuncMap{
			"Quantifier": func(s string) string {
				if strings.Contains(s, "tea") || strings.Contains(s, "coffee") {
					return " bottles of"
				}
				if strings.Contains(s, "cookie") {
					return " bags of"
				}

				return ""
			},
		},
	).Parse(string(b))
	if err != nil {
		return "", err
	}

	buff := bytes.NewBufferString("")

	err = tmpl.Execute(buff, inventory)
	if err != nil {
		return "", err
	}

	return buff.String(), nil
}
