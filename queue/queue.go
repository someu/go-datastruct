package queue

import "go-datastruct/list"

type Queue struct {
	Size int
	list *list.DoublyLinkedList
}

func New() *Queue {
	return &Queue{
		Size: 0,
		list: list.NewDoublyLinkedList(),
	}
}

func (q *Queue) Enqueue(element interface{}) {
	q.list.Insert(q.list.Len, element)
	q.Size += 1
}

func (q *Queue) Dequeue() interface{} {
	node := q.list.Head()
	if node == nil {
		return nil
	}
	q.list.Remove(node)
	q.Size -= 1
	return node.Element
}
