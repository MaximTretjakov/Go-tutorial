package main

import "fmt"

func MapKeys[K comparable, V any](m map[K]V) []K {
	keys := make([]K, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

type List[T any] struct {
	head, tail *element[T]
}

type element[T any] struct {
	next *element[T]
	val  T
}

func (lst *List[T]) Push(v T) {
	if lst.tail == nil {
		lst.head = &element[T]{val: v}
		lst.tail = lst.head
	} else {
		lst.tail.next = &element[T]{val: v}
		lst.tail = lst.tail.next
	}
}

func (lst *List[T]) GetAll() []T {
	var elements []T
	for i := lst.head; i != nil; i = i.next {
		elements = append(elements, i.val)
	}
	return elements
}

func main() {
	m := map[int]string{1: "qwerty", 3: "q", 10: "rom"}
	fmt.Println(MapKeys(m))

	lst := List[int]{}
	lst.Push(10)
	lst.Push(20)
	lst.Push(30)
	fmt.Printf("List elements: %d\n", lst.GetAll())
}
