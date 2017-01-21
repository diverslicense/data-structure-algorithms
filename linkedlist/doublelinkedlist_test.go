package linkedlist

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddToBottomD_NewList(t *testing.T) {
	list := &ListD{}
	list.AddToBottomD(0)
	assert.Equal(t, 1, list.tail.value)
	assert.Equal(t, 0, list.tail.value)
	list.AddToBottomD(1)
	// list.IterateFromHeadD()
	list.IterateFromTailD()
	assert.Equal(t, 1, list.tail.value)
}
