package mapplus

func ToSlice[K comparable, V comparable](m map[K]V) []V {
	var s = make([]V, 0, len(m))
	for _, v := range m {
		s = append(s, v)
	}
	return s
}

func KeyToSlice[K comparable, V comparable](m map[K]V) []K {
	var s = make([]K, 0, len(m))
	for k := range m {
		s = append(s, k)
	}
	return s
}

func Reversal[K comparable, V comparable](m map[K]V) map[V]K {
	var nm = make(map[V]K)
	for k, v := range m {
		nm[v] = k
	}
	return nm
}
