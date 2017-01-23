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
	// TODO: implement with tail
	tail *NodeD
}

// AddToTopD adds the node after the head, at the top of the list
func (Ld *ListD) AddToTopD(value int) {
	new := &NodeD{
		value: value,
	}
	if Ld.head == nil {
		Ld.head = new
	} else {
		tmp := Ld.head
		new.next = tmp
		tmp.prev = new
		Ld.head = new
	}
}

// AddToBottomD adds the node at the end of the list
func (Ld *ListD) AddToBottomD(value int) {
	new := &NodeD{
		value: value,
	}
	if Ld.head == nil {
		Ld.head = new
	} else {
		tmp := Ld.head
		for tmp.next != nil {
			tmp = tmp.next
		}
		tmp.next = new
		new.prev = tmp
	}
}

// RemoveFromTopd removes the node at the top of the list
func (Ld *ListD) RemoveFromTopd() {
	if Ld.head == nil {
		return
	}
	tmp := Ld.head
	if tmp.prev == tmp.next {
		Ld.head = nil
	} else {
		Ld.head = tmp.next
		tmp.next.prev = Ld.head
	}
}

// RemoveFromBottomD removes the node at the end of the list
func (Ld *ListD) RemoveFromBottomD() {
	if Ld.head == nil {
		return
	}
	current := Ld.head
	if current.prev == current.next {
		Ld.head = nil
		return
	}
	for current.next != nil {
		current = current.next
	}
	current.prev.next = nil
}

// IterateFromHeadD iterates through the list from the head and prints out each value
func (Ld *ListD) IterateFromHeadD() {
	current := Ld.head
	for current != nil {
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
