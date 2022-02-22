package stack_test

import (
	"go-datastruct/stack"
	"testing"
)

func TestQueue(t *testing.T) {
	s := stack.New()
	s.Push("A")
	s.Push("B")
	s.Push("C")

	targets := []string{"C", "B", "A"}
	for index, target := range targets {
		curr := s.Top()
		if target != curr {
			t.Errorf("%d top expected to be %s, but got %s", index, target, curr)
		}
		curr = s.Pop()
		if target != curr {
			t.Errorf("%d pop expected to be %s, but got %s", index, target, curr)
		}
	}

	target := s.Top()
	if target != nil {
		t.Errorf("top expected to be nil, but got %s", target)
	}
	target = s.Pop()
	if target != nil {
		t.Errorf("top expected to be nil, but got %s", target)
	}
}
