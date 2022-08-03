package main

import (
	"errors"
	"fmt"
	"testing"
)

func BenchmarkStack_Pop(b *testing.B) {
	stack := NewEmptyStack(5)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		stack.Push(10)

		for item, err := stack.Pop(); err == nil; item, err = stack.Pop() {
			if err != nil {
				fmt.Println(err)
				return
			}

			switch item.(type) {
			case string:
				intItem := (item).(string)
				_ = intItem
			case float64:
				floatItem := item.(float64)
				_ = floatItem
			case Person:
				personItem := (item).(Person)
				_ = personItem.Name
			case int:
				intItem := (item).(int)
				_ = intItem
			default:
				panic(errors.New("unknown type"))
			}
		}
	}
}
