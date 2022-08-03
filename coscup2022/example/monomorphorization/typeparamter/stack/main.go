package main

import (
	"errors"
	"fmt"
)

var (
	ErrEmptySlice = errors.New("pop from empty slice")
)

type Stack[T any] struct {
	items []T
}

func NewStack[T any](cap int) *Stack[T] {
	return &Stack[T]{
		items: make([]T, 0, cap),
	}
}

func (s *Stack[T]) Push(i T) {
	s.items = append(s.items, i)
}

func (s *Stack[T]) Pop() (T, error) {
	if len(s.items) == 0 {
		var zero T
		return zero, ErrEmptySlice
	}

	tmp := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]

	return tmp, nil
}

func main() {
	s := NewStack[int](10)

	for i := 0; i < 10; i++ {
		s.Push(i)
	}

	for i := 0; i < 10; i++ {
		item, err := s.Pop()
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("item + 1 = %d\n", item+1)
	}
}
