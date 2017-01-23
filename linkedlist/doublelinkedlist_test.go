package linkedlist

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddToTop_NewList(t *testing.T) {
	list := &ListD{}
	list.AddToTopD(1)
	list.AddToTopD(2)

	assert.Equal(t, 2, list.head.value)
}

func TestAddToBottomD_NewList(t *testing.T) {
	list := &ListD{}
	list.AddToBottomD(5)
	list.AddToBottomD(6)

	assert.Equal(t, 5, list.head.value)
	assert.Equal(t, 6, list.head.next.value)
}

func TestRemoveFromTopd(t *testing.T) {
	list := &ListD{}
	list.AddToTopD(1)
	list.RemoveFromTopd()

	assert.Nil(t, list.head)
}

func TestRemoveFromBottomd(t *testing.T) {
	list := &ListD{}
	list.AddToBottomD(5)
	list.AddToBottomD(6)
	list.RemoveFromBottomD()

	assert.Nil(t, list.head.next)
	assert.Equal(t, 5, list.head.value)
}
