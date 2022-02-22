package queue_test

import (
	"go-datastruct/queue"
	"testing"
)

func TestQueue(t *testing.T) {
	q := queue.New()
	q.Enqueue("A")
	q.Enqueue("B")
	q.Enqueue("C")

	targets := []string{"A", "B", "C"}
	for index, target := range targets {
		curr := q.Dequeue()
		if target != curr {
			t.Errorf("%d dequeue expected to be %s, but got %s", index, target, curr)
		}
	}

	target := q.Dequeue()
	if target != nil {
		t.Errorf("dequeue expected to be nil, but got %s", target)
	}
}
