package main

import (
	"fmt"
	"reflect"
)

func Filter(data interface{}, fn interface{}) {
	vfn := reflect.ValueOf(fn)
	vdata := reflect.ValueOf(data).Elem()

	j := 0
	for i := 0; i < vdata.Len(); i++ {
		if vfn.Call([]reflect.Value{vdata.Index(i)})[0].Bool() {
			vdata.Index(j).Set(vdata.Index(i)) //
			j++
		}
	}

	vdata.SetLen(j)
}

func main() {
	x := []int{-2, -1, 0, 1, 2, 3, 4}

	Filter(&x, func(n int) bool {
		if n < 0 {
			return false
		}
		return true
	})

	y := []string{"xyz", "zzz", "xxx", "", "x", ""}
	Filter(&y, func(x string) bool {
		if x == "" {
			return false
		}

		return true
	})

	fmt.Println("filter out negative value", x)
	fmt.Println("filter out empty string", y)
}

type myInt int
type myFloat64 float64
