package main

import (
	"test/generator"
)

func main() {
	inventory := generator.Inventory{
		Items: []generator.Item{
			{Name: "black tea", Count: 17},
			{Name: "green tea", Count: 10},
			{Name: "cookie", Count: 5},
			{Name: "coffee", Count: 0},
		},
	}

	generator.Generate(&inventory)
}
