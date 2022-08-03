package main

import (
	"fmt"
	"testing"
)

func BenchmarkStack_Pop(b *testing.B) {
	stack := NewIntEmptyStack(5)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		stack.Push(10)

		for item, err := stack.Pop(); err == nil; item, err = stack.Pop() {
			if err != nil {
				fmt.Println(err)
				return
			}

			_ = item
		}
	}
}
