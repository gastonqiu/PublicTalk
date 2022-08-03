package main

import "testing"

func BenchmarkFilter(b *testing.B) {
	for i := 0; i < b.N; i++ {
		const size = 1000
		x := make([]int, 0, size*2)

		for i := -size; i < size; i++ {
			x = append(x, i)
		}

		b.StartTimer()
		err := Filter(&x, func(i int) bool {
			if i < 0 {
				return true
			}

			return false
		})
		_ = err
		b.StopTimer()
	}
}

//func BenchmarkFilter(b *testing.B) {
//	for i := 0; i < b.N; i++ {
//		const size = 1000
//		x := make([]int, 0, size*2)
//
//		for i := -size; i < size; i++ {
//			x = append(x, i)
//		}
//
//		filter := func(i int) bool {
//			if i < 0 {
//				return false
//			}
//			return true
//		}
//
//		b.StartTimer()
//		FilterInt(&x, filter)
//		b.StopTimer()
//	}
//
//}
