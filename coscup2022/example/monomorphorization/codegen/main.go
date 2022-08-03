package main

import "fmt"

type Person struct {
	ID   int
	Name string
}

func main() {
	stack := NewEmptyIntStack(10)

	for i := 0; i < 10; i++ {
		stack.Push(i)
	}

	for item, err := stack.Pop(); err == nil; item, err = stack.Pop() {
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println(*item)
	}
}
