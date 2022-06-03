package main

import "fmt"

// List represents a singly-linked list that holds
// values of any type.
type List[T any] struct {
	next *List[T]
	val  T
}

func main() {
	root := List[string]{val: "root"}
	one := List[string]{val: "one"}
	// two := List[int]{val: 2}

	// one.next = &two // cannot use &two (value of type *List[int]) as type *List[string] in assignment
	two2 := List[string]{val: "two"}

	one.next = &two2
	root.next = &one

	point := &root
	for point.next != nil {
		fmt.Println(point.val)
		point = point.next
	}
	fmt.Println(point.val)
}
