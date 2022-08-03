package main

import (
	"errors"
)

var (
	ErrEmptySlice = errors.New("pop from empty slice")
)

type StackString struct {
	items []string
}

func NewEmptyStringStack(cap int) *StackString {
	return &StackString{
		items: make([]string, 0, cap),
	}
}

func (s *StackString) Push(i string) {
	s.items = append(s.items, i)
}

func (s *StackString) Pop() (*string, error) {
	if len(s.items) == 0 {
		return nil, ErrEmptySlice
	}

	tmp := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]

	return &tmp, nil
}

type StackInt struct {
	items []int
}

func NewEmptyIntStack(cap int) *StackInt {
	return &StackInt{
		items: make([]int, 0, cap),
	}
}

func (s *StackInt) Push(i int) {
	s.items = append(s.items, i)
}

func (s *StackInt) Pop() (*int, error) {
	if len(s.items) == 0 {
		return nil, ErrEmptySlice
	}

	tmp := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]

	return &tmp, nil
}

type StackPerson struct {
	items []Person
}

func NewEmptyPersonStack(cap int) *StackPerson {
	return &StackPerson{
		items: make([]Person, 0, cap),
	}
}

func (s *StackPerson) Push(i Person) {
	s.items = append(s.items, i)
}

func (s *StackPerson) Pop() (*Person, error) {
	if len(s.items) == 0 {
		return nil, ErrEmptySlice
	}

	tmp := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]

	return &tmp, nil
}
