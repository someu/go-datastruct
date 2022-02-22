package list

type CycleLinkedNode struct {
	Element interface{}
	Next    *CycleLinkedNode
	Prev    *CycleLinkedNode
}

type CycleLinkedList struct {
	head *CycleLinkedNode
	Len  int
}

func NewCycleLinkedList() *CycleLinkedList {
	return &CycleLinkedList{
		head: nil,
		Len:  0,
	}
}

func (list *CycleLinkedList) Insert(loc int, element interface{}) *CycleLinkedNode {
	if loc > list.Len {
		return nil
	}
	nextNode := list.Get(loc)
	newNode := &CycleLinkedNode{
		Element: element,
		Next:    nil,
		Prev:    nil,
	}
	if nextNode == nil {
		list.head = newNode
		newNode.Next = newNode
		newNode.Prev = newNode
	} else {
		prevNode := nextNode.Prev
		newNode.Prev = prevNode
		prevNode.Next = newNode
		newNode.Next = nextNode
		nextNode.Prev = newNode
		if loc == 0 {
			list.head = newNode
		}
	}
	list.Len += 1
	return newNode
}

func (list *CycleLinkedList) Remove(node *CycleLinkedNode) {
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
	prevNode := node.Prev
	nextNode := node.Next

	if list.Len == 1 {
		list.head = nil
	} else {
		prevNode.Next = nextNode
		nextNode.Prev = prevNode
		if list.head == node {
			list.head = nextNode
		}
	}

	list.Len -= 1
}

func (list *CycleLinkedList) Get(loc int) *CycleLinkedNode {
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

func (list *CycleLinkedList) Find(element interface{}) (int, *CycleLinkedNode) {
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
