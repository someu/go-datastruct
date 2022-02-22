package queue

import "go-datastruct/list"

type Queue struct {
	list *list.DoublyLinkedList
}

func New() *Queue {
	return &Queue{
		list: list.NewDoublyLinkedList(),
	}
}

func (q *Queue) Enqueue(element interface{}) {
	q.list.Insert(q.list.Len, element)
}

func (q *Queue) Dequeue() interface{} {
	node := q.list.Head()
	if node == nil {
		return nil
	}
	q.list.Remove(node)
	return node.Element
}
