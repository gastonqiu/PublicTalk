package main

import (
	"fmt"
	"os"
	"text/template"
)

func main() {
	init := `package main

import (
	"errors"
)

var (
	ErrEmptySlice = errors.New("pop from empty slice")
)
`
	stackTemplate := fmt.Sprint(`

type Stack{{.FuncName}} struct {
	items []{{.TypeName}}
}

func NewEmpty{{.FuncName}}Stack(cap int) *Stack{{.FuncName}} {
	return &Stack{{.FuncName}}{
		items: make([]{{.TypeName}}, 0, cap),
	}
}

func (s *Stack{{.FuncName}}) Push(i {{.TypeName}}) {
	s.items = append(s.items, i)
}

func (s *Stack{{.FuncName}}) Pop() (*{{.TypeName}}, error) {
	if len(s.items) == 0 {
		return nil, ErrEmptySlice
	}

	tmp := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]

	return &tmp, nil
}

`)

	file, err := os.Create("../stack.go")
	if err != nil {
		panic(err)
	}

	_, err = file.WriteString(init)
	if err != nil {
		panic(err)
	}

	tmpl, err := template.New("test").Parse(stackTemplate)
	if err != nil {
		panic(err)
	}

	typeStr := []struct {
		TypeName string
		FuncName string
	}{
		{
			TypeName: "string",
			FuncName: "String",
		},
		{
			TypeName: "int",
			FuncName: "Int",
		},
		{
			TypeName: "Person",
			FuncName: "Person",
		},
	}

	for _, t := range typeStr {

		err = tmpl.Execute(file, t)
		if err != nil {
			panic(err)
		}
	}
}
