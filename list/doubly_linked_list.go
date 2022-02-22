package list

type DoublyLinkedNode struct {
	Element interface{}
	Next    *DoublyLinkedNode
	Prev    *DoublyLinkedNode
}

type DoublyLinkedList struct {
	head *DoublyLinkedNode
	tail *DoublyLinkedNode
	Len  int
}

func NewDoublyLinkedList() *DoublyLinkedList {
	return &DoublyLinkedList{
		head: nil,
		tail: nil,
		Len:  0,
	}
}

func (list *DoublyLinkedList) Insert(loc int, element interface{}) *DoublyLinkedNode {
	if loc > list.Len {
		return nil
	}
	nextNode := list.Get(loc)
	newNode := &DoublyLinkedNode{
		Element: element,
		Next:    nextNode,
		Prev:    nil,
	}
	if loc == 0 && nextNode == nil {
		list.head = newNode
		list.tail = newNode
	} else if loc == 0 && nextNode != nil {
		list.head = newNode
		nextNode.Prev = newNode
	} else if loc != 0 && nextNode == nil {
		newNode.Prev = list.tail
		list.tail.Next = newNode
		list.tail = newNode
	} else {
		nextNode.Prev.Next = newNode
		newNode.Prev = nextNode.Prev
		nextNode.Prev = newNode
	}
	list.Len += 1
	return newNode
}

func (list *DoublyLinkedList) Remove(node *DoublyLinkedNode) {
	len := list.Len
	curr := list.head
	for {
		if len <= 0 {
			return
		}
		if curr == node {
			break
		}
		curr = curr.Next
		len--
	}
	prev := node.Prev
	next := node.Next
	if prev == nil && next == nil {
		list.head = nil
		list.tail = nil
	} else if prev == nil && next != nil {
		list.head = next
		next.Prev = nil
	} else if prev != nil && next == nil {
		list.tail = prev
		prev.Next = nil
	} else {
		prev.Next = next
		next.Prev = prev
	}
	list.Len -= 1
}

func (list *DoublyLinkedList) Get(loc int) *DoublyLinkedNode {
	if loc >= list.Len {
		return nil
	}

	curr := list.head
	for {
		if loc == 0 {
			return curr
		}
		curr = curr.Next
		loc -= 1
	}
}

func (list *DoublyLinkedList) Find(element interface{}) (int, *DoublyLinkedNode) {
	curr := list.head
	index := 0
	for {
		if index >= list.Len {
			return -1, nil
		}
		if curr.Element == element {
			return index, curr
		}
		curr = curr.Next
		index++
	}
}
