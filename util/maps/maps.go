package maps

func GetKeys[M ~map[K]V, K comparable, V any](m M) []K {
	data := make([]K, 0, len(m))
	for k := range m {
		data = append(data, k)
	}
	return data
}
