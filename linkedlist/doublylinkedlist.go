package linkedlist

import "fmt"

// NodeD is a node of a doubly linked list
type NodeD struct {
	prev  *NodeD
	value int
	next  *NodeD
}

// ListD contains pointers to the head and tail of the list
type ListD struct {
	head *NodeD
	tail *NodeD
}

// AddToBottomD adds the node before the tail, at the end of the list
func (Ld *ListD) AddToBottomD(value int) {
	new := &NodeD{
		value: value,
	}
	if Ld.tail == nil {
		Ld.tail = new
		new.prev = Ld.head
		Ld.head.next = Ld.tail
	} else {
		new.prev = Ld.tail.prev
		Ld.tail = new
	}
}

// AddToTopD adds the node after the head, at the top of the list
func (Ld *ListD) AddToTopD(value int) {
	new := &NodeD{
		value: value,
	}
	if Ld.head.next != nil {
		new.next = Ld.head.next
		new.prev = Ld.head
	}
	Ld.head.next = new
}

// RemoveFromBottomD removes the node at the end of the list
func (Ld *ListD) RemoveFromBottomD() {
	Ld.tail.prev = Ld.tail.prev.prev
	Ld.tail.prev.next = Ld.tail
}

// RemoveFromTopd removes the node at the top of the list
func (Ld *ListD) RemoveFromTopd() {
	Ld.head.next = Ld.head.next.next
	Ld.head.next.prev = Ld.head
}

// IterateFromHeadD iterates through the list from the head and prints out each value
func (Ld *ListD) IterateFromHeadD() {
	current := Ld.head
	for current.next != nil {
		fmt.Printf("%d ", current.value)
		current = current.next
	}
}

// IterateFromTailD iterates through the list from the tail and prints out each value
func (Ld *ListD) IterateFromTailD() {
	current := Ld.tail
	for current.prev != nil {
		fmt.Printf("%d ", current.value)
		current = current.prev
	}
}
