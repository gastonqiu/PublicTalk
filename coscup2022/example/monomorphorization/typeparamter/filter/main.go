package main

import "fmt"

func Filter[T any](data *[]T, fn func(n T) bool) {
	j := 0
	for idx, d := range *data {
		if fn(d) {
			(*data)[j] = (*data)[idx]
			j++
		}
	}

	*data = (*data)[:j]
}

func main() {
	const min = -10
	const max = 10
	x := make([]int, 0, max-min)
	for i := min; i < max; i++ {
		x = append(x, i)
	}

	Filter[int](&x, func(x int) bool {
		if x < 0 {
			return false
		}
		return true
	})

	fmt.Printf("result: %v\n", x)
}
