package main

import "fmt"

// List represents a singly-linked list that holds
// values of any type.
type List[T any] struct {
	next *List[T]
	val  T
}

func (list *List[T]) add(e T) {
	n := list
	for n.next != nil {
		n = n.next
	}
	n.next = &List[T]{nil, e}
}

func (list *List[T]) toSlice() []T {
	s := make([]T, 0)
	n := list
	for n != nil {
		s = append(s, n.val)
		n = n.next
	}
	return s
}

func main() {
	list := List[string]{nil, "hello"}
	list.add("world")
	s := list.toSlice()
	fmt.Println(s)
	numlist := List[int]{nil, 42}
	numlist.add(999)
	t := numlist.toSlice()
	fmt.Println(t)
}
