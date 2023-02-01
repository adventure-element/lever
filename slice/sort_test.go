package slice

import (
	"fmt"
	"testing"
)

func TestInsert(t *testing.T) {
	var s []int = []int{0, 2, 3, 4, 5, 7}
	Insert(&s, 1, 1)
	Insert(&s, 5, 6)
	fmt.Println(s)
}

func TestDel(t *testing.T) {
	var s = []int{0, 1, 2, 3, 4, 5}
	s = append(s, 6)
	fmt.Println(s)
	Del(&s, 2)
	fmt.Println(s)
}

func TestMove(t *testing.T) {
	var s = []int{0, 1, 2, 3, 4, 5, 6, 7}
	Move(&s, 2, 4)
	fmt.Println(s)
}
