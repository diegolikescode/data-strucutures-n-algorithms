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
	list.Append(4200)

	assert.NotEqual(t, list, nil, "List must be initialized")
	assert.Equal(t, 2, list.Length, "Must be able to append item")
}

func testPopList(t *testing.T) {
	list := InitSinglyLinkedList[int]()

	list.Append(69)
	list.Append(4200)

	assert.NotEqual(t, list, nil, "List must be initialized")
	assert.Equal(t, 2, list.Length, "Must be able to append item")

	assert.Equal(t, 69, list.Pop().Data, "Head of list must be correct when list is Popped")
	assert.Equal(t, 1, list.Length, "Must be able to append item")

	assert.Equal(t, 4200, list.Pop().Data, "Head of list must be correct when list is Popped")
	assert.Equal(t, 0, list.Length, "Must be able to append item")
}

func testSearchData(t *testing.T) {
	list := InitSinglyLinkedList[int]()

	list.Append(1)
	list.Append(2)
	list.Append(3)
	list.Append(4)

	assert.Equal(t, 4, list.Length, "List length must be 4")

	found := 4
	n := list.SearchData(found)
	assert.NotNil(t, n, "Node should have been found (not nil)")
	assert.Equal(t, found, n.Data, "numbers must be equal of the search")

	notFound := 10
	notN := list.SearchData(notFound)
	assert.Nil(t, notN, "Node should have been found (not nil)")
}

func TestSinglyLinkedList(t *testing.T) {
	testInitializeList(t)
	testAppendToList(t)
	testPopList(t)
	testSearchData(t)
}
