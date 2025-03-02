package datastructure

type Node[T any] struct {
	Data T
}

type SinglyLinkedList struct {
	Head   *Node[any]
	Tail   *Node[any]
	Length int
}

func (s *SinglyLinkedList) Append() {
}

func (s *SinglyLinkedList) Prepend() {
}

func (s *SinglyLinkedList) Pop() *Node[any] {
	return nil
}

func (s *SinglyLinkedList) Search() {
}

func InitSinglyLinkedList() SinglyLinkedList {
	return SinglyLinkedList{
		Head:   nil,
		Tail:   nil,
		Length: 0,
	}
}
