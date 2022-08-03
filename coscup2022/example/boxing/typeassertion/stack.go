package main

import (
	"errors"
	"fmt"
)

var (
	ErrEmptySlice = errors.New("pop from empty slice")
)

type Stack struct {
	items []interface{}
}

func NewEmptyStack(cap int) *Stack {
	return &Stack{
		items: make([]interface{}, 0, cap),
	}
}

func (s *Stack) Push(i interface{}) {
	s.items = append(s.items, i)
}

func (s *Stack) Pop() (interface{}, error) {
	if len(s.items) == 0 {
		return 0, ErrEmptySlice
	}

	tmp := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]

	return tmp, nil
}

type Person struct {
	ID   int
	Name string
}

func main() {
	stack := NewEmptyStack(5)

	p1 := Person{ID: 1, Name: "Gaston"}
	stack.Push(p1)

	for item, err := stack.Pop(); err == nil; item, err = stack.Pop() {
		if err != nil {
			fmt.Println(err)
			return
		}

		switch item.(type) {
		case Person:
			personItem, OK := (item).(Person)
			fmt.Printf("name: %s, id: %d, ok: %v", personItem.Name, personItem.ID, OK)
		case int:
			intItem := (item).(int)
			fmt.Printf("int item: %d", intItem)
		default:
			panic(errors.New("unknown type"))
		}
	}
}
