package list

type LinkedNode struct {
	Element interface{}
	Next    *LinkedNode
}

type LinkedList struct {
	head *LinkedNode
	Len  int
}

func NewLinkedList() *LinkedList {
	return &LinkedList{}
}

func (list *LinkedList) Insert(loc int, element interface{}) *LinkedNode {
	if loc > list.Len {
		return nil
	}
	curr := list.head
	var prev *LinkedNode
	for {
		if loc == 0 {
			node := &LinkedNode{
				Element: element,
				Next:    curr,
			}
			if prev == nil {
				list.head = node
			} else {
				prev.Next = node
			}
			list.Len += 1
			return node
		}
		prev = curr
		curr = curr.Next
		loc -= 1
	}
}

func (list *LinkedList) Remove(node *LinkedNode) {
	var prev *LinkedNode
	curr := list.head
	for {
		if curr == nil {
			return
		}
		if curr == node {
			if prev == nil {
				list.head = curr.Next
			} else {
				prev.Next = curr.Next
			}
			list.Len -= 1
			return
		}
		prev = curr
		curr = curr.Next
	}
}

func (list *LinkedList) Get(loc int) *LinkedNode {
	if loc >= list.Len {
		return nil
	}
	curr := list.head

	for {
		if loc == 0 {
			return curr
		}
		curr = curr.Next
		loc--
	}
}

func (list *LinkedList) Find(element interface{}) (int, *LinkedNode) {
	curr := list.head
	loc := 0
	for {
		if curr == nil {
			return -1, nil
		}
		if curr.Element == element {
			return loc, curr
		}
		curr = curr.Next
		loc += 1
	}
}

func (list *LinkedList) Head() *LinkedNode {
	return list.head
}
