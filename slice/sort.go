package slice

func Del[V any](slice *[]V, index int) {
	s := *slice
	*slice = append(s[:index], s[index+1:]...)
}

func Insert[V any](slice *[]V, index int, value V) {
	s := *slice
	if index <= 0 {
		*slice = append([]V{value}, s...)
	} else if index >= len(s) {
		*slice = append(s, value)
	} else {
		*slice = append(s[:index], append([]V{value}, s[index:]...)...)
	}
}

func Move[V any](slice *[]V, index, to int) {
	s := *slice
	v := s[index]
	if index == to {
		return
	} else if to < index {
		Del[V](slice, index)
		Insert(slice, to, v)
	} else {
		Insert(slice, to, v)
		Del[V](slice, index)
	}
}
