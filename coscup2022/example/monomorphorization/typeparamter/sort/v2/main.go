package main

import (
	"fmt"
	"sort"
)

type orderedSlice[T any] struct {
	s   []T
	cmp func(T, T) bool
}

func (s orderedSlice[T]) Len() int {
	return len(s.s)
}

func (s orderedSlice[T]) Less(i, j int) bool {
	return s.cmp(s.s[i], s.s[j])
}

func (s orderedSlice[T]) Swap(i, j int) {
	s.s[i], s.s[j] = s.s[j], s.s[i]
}

// OrderedSlice sorts the slice s in ascending order.
// The elements of s must be ordered using the < operator.
func OrderedSlice[T any](s []T, cmp func(T, T) bool) {
	// Convert s to the type orderedSlice[T].
	// As s is []T, and orderedSlice[T] is defined as []T,
	// this conversion is permitted.
	// orderedSlice[T] implements sort.Interface,
	// so can pass the result to sort.Sort.
	// The elements will be sorted using the < operator.
	orderS := &orderedSlice[T]{
		s:   s,
		cmp: cmp,
	}

	sort.Sort(orderS)
}

func main() {
	x := []int{3, 2, 1}
	OrderedSlice[int](x, func(a int, b int) bool {
		return a < b
	})
	fmt.Println(x)

	type Person struct {
		Name string
	}

	p := []Person{
		{Name: "Henry"},
		{Name: "Gaston"},
		{Name: "Amy"},
	}
	OrderedSlice[Person](p, func(p1 Person, p2 Person) bool {
		return p1.Name < p2.Name
	})
}
