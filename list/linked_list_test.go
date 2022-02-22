package list_test

import (
	"go-datastruct/list"
	"testing"
)

func TestLinkedList(t *testing.T) {
	l := list.NewLinkedList()
	l.Insert(0, "A")  // A
	l.Insert(0, "B")  // B A
	l.Insert(0, "C")  // C B A
	l.Insert(1, "D")  // C D B A
	l.Insert(10, "F") // C D B A
	l.Insert(1, "G")  // C G D B A
	l.Insert(2, "H")  // C G H D B A
	l.Insert(3, "I")  // C G H I D B A

	l.Remove(&list.LinkedNode{})
	l.Remove(l.Get(0)) // G H I D B A
	l.Remove(l.Get(1)) // G I D B A

	if l.Get(10) != nil {
		t.Errorf("node 10 expected to be nil")
	}

	iA, nA := l.Find("A")
	if iA == 4 && nA == nil {
		t.Errorf("node A expected to A, but got nil")
	} else if nA.Element != "A" {
		t.Errorf("node A expected to be A:4, bug got %s:%d", nA.Element, iA)
	}
	if _, nM := l.Find("M"); nM != nil {
		t.Errorf("node M expected to be nil")
	}

	targets := []string{"G", "I", "D", "B", "A"}
	for index, target := range targets {
		n := l.Get(index)
		if n == nil {
			t.Errorf("node %d expected to be %s, but got nil", index, target)
		} else {
			if n.Element != target {
				t.Errorf("node %d expected to be %s, but got %s", index, target, n.Element)
			}
			if index != len(targets)-1 {
				next := targets[index+1]
				if n.Next == nil {
					t.Errorf("node %d's next expected to be %s, but got nil", index, next)
				} else if n.Next.Element != next {
					t.Errorf("node %d's next expected to be %s, but got %s", index, next, n.Next.Element)
				}
			}
		}
	}
}
