package stack

import "go-datastruct/list"

type Stack struct {
	Size int
	list *list.DoublyLinkedList
}

func New() *Stack {
	return &Stack{
		Size: 0,
		list: list.NewDoublyLinkedList(),
	}
}

func (s *Stack) Push(element interface{}) {
	s.list.Insert(s.list.Len, element)
	s.Size += 1
}

func (s *Stack) Pop() interface{} {
	node := s.list.Tail()
	if node == nil {
		return nil
	}
	s.list.Remove(node)
	s.Size -= 1
	return node.Element
}

func (s *Stack) Top() interface{} {
	tail := s.list.Tail()
	if tail == nil {
		return nil
	}
	return tail.Element
}
