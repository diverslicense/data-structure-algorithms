package linkedlist

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddToBottom_NewList_CreateNewList(t *testing.T) {
	l := List{}
	l.AddToBottom(1)
	l.AddToBottom(2)

	assert.Equal(t, 1, l.head.value)
	assert.Equal(t, 2, l.head.next.value)
}

func TestAddToTop_NewList_CreateNewList(t *testing.T) {
	l := List{}
	l.AddToTop(0)
	l.AddToTop(1)

	assert.Equal(t, 1, l.head.value)
	assert.Equal(t, 0, l.head.next.value)
}

func TestRemoveFromBottom(t *testing.T) {
	l := List{}
	l.AddToBottom(1)
	l.AddToBottom(2)
	l.RemoveFromBottom()

	assert.Equal(t, 1, l.head.value)
	assert.Nil(t, l.head.next)
}

func TestRemoveFromTop(t *testing.T) {
	l := List{}
	l.AddToTop(1)
	l.AddToTop(0)
	l.RemoveFromTop()

	assert.Equal(t, 1, l.head.value)
}
