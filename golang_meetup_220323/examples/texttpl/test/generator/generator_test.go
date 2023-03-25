package generator

import (
	"os"
	"test/strcmp"
	"testing"
)

func TestGenerate(t *testing.T) {
	inventory := Inventory{
		Items: []Item{
			{"black tea", 17},
			{"green tea", 10},
			{"cookie", 5},
			{"coffee", 1},
		},
	}

	s, err := Generate(&inventory)
	if err != nil {
		t.Error(err)
	}

	b, err := os.ReadFile("../_test/text")
	if err != nil {
		t.Error(err)
	}

	if string(b) != s {
		t.Errorf("expected: %s, actual: %s", b, s)
	}
}

func TestGenerateErr(t *testing.T) {
	inventory := Inventory{
		Items: []Item{
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

	b, err := os.ReadFile("../_test/text_2")
	if err != nil {
		t.Error(err)
	}

	bs := string(b)

	if bs != s {
		diff := strcmp.ANSIDiff(bs, s)
		t.Errorf("file: _test/text\ndiff:\n%s", diff)
	}
}
