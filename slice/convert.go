package slice

import "github.com/adventure-element/lever/mapplus"

func ToMap[V comparable](slice []V) map[int]V {
	var m = make(map[int]V)
	for i, v := range slice {
		m[i] = v
	}
	return m
}

func ToSortMap[V comparable](slice []V) mapplus.Sort[int, V] {
	var m mapplus.Sort[int, V]
	for i, v := range slice {
		m.Set(i, v)
	}
	return m
}
