package linkedlist

import "fmt"

// 7. Singly Linked List: Add or Remove at top or bottom of list
type Node struct {
	value int
	next  *Node
}

type List struct {
	head *Node
}

func (L *List) AddToBottom(value int) {
	newItem := &Node{
		value: value,
	}
	// first check if list exists
	if L.head == nil {
		L.head = newItem
	} else {
		// iterate through to last item in list
		current := L.head
		for current.next != nil {
			current = current.next
		}
		current.next = newItem
	}
}

func (L *List) AddToTop(value int) {
	item := &Node{
		value: value,
		next:  L.head,
	}
	L.head = item
}

func (L *List) RemoveFromBottom() {
	if L.head != nil {
		current := L.head
		for current.next.next != nil {
			current = current.next
		}
		current.next = current.next.next
	}
}

func (L *List) RemoveFromTop() {
	if L.head != nil {
		L.head = L.head.next
	}
}

func (L *List) Iterate() {
	current := L.head
	for current != nil {
		fmt.Printf("%v ", current.value)
		current = current.next
	}
	fmt.Println()
}
