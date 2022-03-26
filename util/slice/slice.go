package slice

func Fill[T any](size int, genf func() T) []T {
	data := make([]T, 0, size)
	for i := 0; i < cap(data); i++ {
		data = append(data, genf())
	}
	return data
}

func ToMap[T ~[]I, I any, K comparable](xs T, op func(I) K) map[K]I {
	data := make(map[K]I, len(xs))
	for i := range xs {
		data[op(xs[i])] = xs[i]
	}
	return data
}

func Contains[T ~[]I, I comparable](haystack T, needle I) bool {
	for i := range haystack {
		if needle == haystack[i] {
			return true
		}
	}
	return false
}
