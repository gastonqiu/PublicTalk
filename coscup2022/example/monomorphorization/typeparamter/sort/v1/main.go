package main

import (
	"fmt"
	"sort"
)

// Ordered is a type constraint that matches all ordered types.
// (An ordered type is one that supports the < <= >= > operators.)
// In practice this type constraint would likely be defined in
// a standard library package.
type Ordered interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
	~float32 | ~float64 |
	~string
}

type orderedSlice[T Ordered] struct {
	s []T
}

func (s orderedSlice[T]) Len() int {
	return len(s.s)
}

func (s orderedSlice[T]) Less(i, j int) bool {
	return s.s[i] < s.s[j]
}

func (s orderedSlice[T]) Swap(i, j int) {
	s.s[i], s.s[j] = s.s[j], s.s[i]
}

func OrderedSlice[T Ordered](s []T) {
	orderS := &orderedSlice[T]{
		s: s,
	}

	sort.Sort(orderS)
}

func main() {
	x := []int{3, 2, 1}
	OrderedSlice[int](x)
	fmt.Println(x)

	y := []string{"d", "c", "a", "b"}
	OrderedSlice[string](y)
	fmt.Println(y)
}
