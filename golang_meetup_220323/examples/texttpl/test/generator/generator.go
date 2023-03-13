package generator

import (
	"bytes"
	"html/template"
	"os"
)

type Inventory struct {
	Item []Item
}

type Item struct {
	Name  string
	Count uint
}

func Generate(inventory *Inventory) (string, error) {
	b, err := os.ReadFile("./text.tpl")
	if err != nil {
		return "", err
	}

	tmpl, err := template.New("test").Parse(string(b))
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
