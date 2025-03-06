package datastructure

type Node[T comparable] struct {
	Data T
	Next *Node[T]
}

type SinglyLinkedList[T comparable] struct {
	Head   *Node[T]
	Length int
}

func (s *SinglyLinkedList[T]) Append(data T) {
	newNode := InitNode(data)

	if s.Length == 0 {
		s.Head = newNode
		s.Length++
		return
	}

	curr := s.Head
	for curr.Next != nil {
		curr = curr.Next
	}
	curr.Next = newNode
	s.Length++
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

func (s *SinglyLinkedList[T]) SearchData(data T) *Node[T] {
	if s.Length <= 0 {
		return nil
	}

	curr := s.Head
	for curr != nil {
		if curr.Data == data {
			return curr
		}
		curr = curr.Next
	}
	return nil
}

func InitNode[T comparable](data T) *Node[T] {
	return &Node[T]{Data: data, Next: nil}
}

func InitSinglyLinkedList[T comparable]() SinglyLinkedList[T] {
	return SinglyLinkedList[T]{
		Head:   nil,
		Length: 0,
	}
}
