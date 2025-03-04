package datastructure

type Node[T any] struct {
	Data T
	Next *Node[T]
}

type SinglyLinkedList[T any] struct {
	Head   *Node[T]
	Length int
}

func (s *SinglyLinkedList[T]) Append(data T) {
	newNode := Node[T]{
		Data: data,
		Next: nil,
	}

	if s.Length == 0 {
		s.Head = &newNode
		s.Length++
		return
	}

	curr := *s.Head
	for curr.Next != nil {
		curr = *curr.Next
	}
	curr.Next = &newNode
}

func (s *SinglyLinkedList[T]) Prepend() {
}

func (s *SinglyLinkedList[T]) Pop() *Node[T] {
	if s.Head == nil {
		return nil
	}

	node := s.Head
	s.Head = s.Head.Next
	s.Length--

	return node
}

func (s *SinglyLinkedList[T]) Search() {
}

func InitNode[T any](data T) *Node[T] {
	return &Node[T]{Data: data, Next: nil}
}

func InitSinglyLinkedList[T any]() SinglyLinkedList[T] {
	return SinglyLinkedList[T]{
		Head:   nil,
		Length: 0,
	}
}
