package main

import (
	"fmt"
	"reflect"
)

func Filter(data interface{}, fn interface{}) error {
	// validation input data is pointer of slice
	if reflect.ValueOf(data).Kind() != reflect.Pointer {
		return fmt.Errorf("reflection: data is not a slice")
	}
	vdata := reflect.ValueOf(data).Elem()
	if vdata.Kind() != reflect.Slice {
		return fmt.Errorf("reflection: data is not a slice")
	}

	// validate input data is func type
	vfn := reflect.ValueOf(fn)
	vfnKind := vfn.Kind()
	vfnType := vfn.Type()
	if vfnKind != reflect.Func {
		return fmt.Errorf("reflection: fn is not func type")
	}

	// check the input func is type of func(i genericType) bool
	if vfnType.NumIn() != 1 || vfnType.In(0).Kind() != vdata.Index(0).Kind() {
		return fmt.Errorf("reflection: fn input params doesn't match")
	}

	if vfnType.NumOut() != 1 || vfnType.Out(0).Kind() != reflect.Bool {
		return fmt.Errorf("reflection: fn output params doesn't match")
	}

	// filter
	j := 0
	for i := 0; i < vdata.Len(); i++ {
		if vfn.Call([]reflect.Value{vdata.Index(i)})[0].Bool() {
			vdata.Index(j).Set(vdata.Index(i))
			j++
		}
	}

	vdata.SetLen(j)

	return nil
}

func FilterInt(data *[]int, fn func(n int) bool) {
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
	x := []int{-2, -1, 0, 1, 2, 3, 4}

	err := Filter(&x, func(n int) bool {
		if n < 0 {
			return false
		}
		return true
	})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("--------------------------------")

	fmt.Printf("validate case: %v\n", x)

	fmt.Println("--------------------------------")

	x = []int{-2, -1, 0, 1, 2, 3, 4}
	err = Filter(&x, func() bool {
		return true
	})
	if err != nil {
		fmt.Printf("invalidate case: %v\n", err)
	}

}
