package main

import (
	"errors"
	"fmt"
)

var (
	ErrEmptySlice = errors.New("pop from empty slice")
)

type StackInt struct {
	items []int
}

func NewIntEmptyStack(cap int) *StackInt {
	return &StackInt{
		items: make([]int, 0, cap),
	}
}

func (s *StackInt) Push(i int) {
	s.items = append(s.items, i)
}

func (s *StackInt) Pop() (int, error) {
	if len(s.items) == 0 {
		return 0, ErrEmptySlice
	}

	tmp := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]

	return tmp, nil
}

type StackStr struct {
	items []string
}

func NewStringEmptyStack(cap int) *StackStr {
	return &StackStr{
		items: make([]string, 0, cap),
	}
}

func (s *StackStr) Push(i string) {
	s.items = append(s.items, i)
}

func (s *StackStr) Pop() (string, error) {
	if len(s.items) == 0 {
		return "", ErrEmptySlice
	}

	tmp := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]

	return tmp, nil
}

func main() {
	strStack := NewStringEmptyStack(5)
	strStack.Push("H")
	strStack.Push("e")
	strStack.Push("l")
	strStack.Push("l")
	strStack.Push("o")

	fmt.Print("str stack: ")

	for item, err := strStack.Pop(); err == nil; item, err = strStack.Pop() {
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Print(item)
	}
	fmt.Print("\n")

	fmt.Print("int stack: ")

	intStack := NewIntEmptyStack(5)
	intStack.Push(1)
	intStack.Push(2)
	intStack.Push(3)
	intStack.Push(4)
	intStack.Push(5)

	for item, err := intStack.Pop(); err == nil; item, err = intStack.Pop() {
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Print(item)
	}

}
