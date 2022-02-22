package list_test

import (
	"go-datastruct/list"
	"testing"
)

func TestCycleLinkedList(t *testing.T) {
	l := list.NewCycleLinkedList()
	l.Insert(0, "A")  // A
	l.Insert(0, "B")  // B A
	l.Insert(0, "C")  // C B A
	l.Insert(1, "D")  // C D B A
	l.Insert(10, "F") // C D B A
	l.Insert(1, "G")  // C G D B A
	l.Insert(2, "H")  // C G H D B A
	l.Insert(3, "I")  // C G H I D B A

	l.Remove(&list.CycleLinkedNode{})
	l.Remove(l.Get(0)) // G H I D B A
	l.Remove(l.Get(1)) // G I D B A

	if l.Get(10) != nil {
		t.Errorf("node 10 expected to be nil")
	}

	iA, nA := l.Find("A")
	if nA == nil {
		t.Errorf("node A expected to A, but got nil")
	} else if nA.Element != "A" || iA != 4 {
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
			expectPrev := ""
			expectNext := ""
			prev := ""
			next := ""

			if index != 0 {
				expectPrev = targets[index-1]
			} else {
				expectPrev = targets[len(targets)-1]
			}
			if index != len(targets)-1 {
				expectNext = targets[index+1]
			} else {
				expectNext = targets[0]
			}
			if n.Prev != nil {
				prev = n.Prev.Element.(string)
			}
			if n.Next != nil {
				next = n.Next.Element.(string)
			}

			if expectNext != next {
				t.Errorf("node %d's next expected to be %s, but got %s", index, expectNext, next)
			}
			if expectPrev != prev {
				t.Errorf("node %d's prev expected to be %s, but got %s", index, expectPrev, prev)
			}
		}
	}

	head := l.Head()
	if nA == nil {
		t.Errorf("head expected to G, but got nil")
	} else if head.Element != "G" {
		t.Errorf("head expected to be G, bug got %s", head.Element)
	}

	l.Remove(l.Get(4))
	l.Remove(l.Get(3))
	l.Remove(l.Get(2))
	l.Remove(l.Get(1))
	l.Remove(l.Get(0))
}
