package util

func FillSlice[T any](size int, genf func() T) []T {
	data := make([]T, 0, size)
	for i := 0; i < cap(data); i++ {
		data = append(data, genf())
	}
	return data
}

func SemiCall1[A any, R any](f func(A) R, a A) func() R {
	return func() R {
		return f(a)
	}
}

func SliceToMap[T ~[]I, I any, K comparable](xs T, op func(I) K) map[K]I {
	data := make(map[K]I, len(xs))
	for i := range xs {
		data[op(xs[i])] = xs[i]
	}
	return data
}
