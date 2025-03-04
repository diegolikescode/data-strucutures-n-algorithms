package datastructure

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func testInitializeList(t *testing.T) {
	list := InitSinglyLinkedList[int]()

	assert.NotEqual(t, list, nil, "Should have a initialized list")
	assert.Equal(t, list.Length, 0, "Should have a initialized list")
}

func testAppendToList(t *testing.T) {
	list := InitSinglyLinkedList[int]()

	list.Append(69)

	assert.NotEqual(t, list, nil, "List must be initialized")
	assert.Equal(t, list.Length, 1, "List must be initialized")
}

func TestSinglyLinkedList(t *testing.T) {
	testInitializeList(t)
	testAppendToList(t)
}
