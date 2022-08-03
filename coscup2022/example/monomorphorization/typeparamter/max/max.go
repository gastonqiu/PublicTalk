package max

type Ordered interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
		~float32 | ~float64 |
		~string
}

type ss string

func min[T Ordered](a T, b T) T {
	if a < b {
		return a
	}

	return b
}

func MinInt(a int, b int) int {
	return min[int](a, b)
}

func MinStr(a string, b string) string {
	return min[string](a, b)
}

func MinStr2(a ss, b ss) ss {
	return min[ss](a, b)
}
