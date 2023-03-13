package generator

import (
	"os"
	"testing"
)

func TestGenerate(t *testing.T) {
	inventory := Inventory{
		Item: []Item{
			{"black tea", 17},
			{"green tea", 10},
			{"cookie", 5},
			{"coffee", 0},
		},
	}

	s, err := Generate(&inventory)
	if err != nil {
		t.Error(err)
	}

	b, err := os.ReadFile("text")
	if err != nil {
		t.Error(err)
	}

	if string(b) != s {
		t.Errorf("expected: %s, actual: %s", b, s)
	}
}
