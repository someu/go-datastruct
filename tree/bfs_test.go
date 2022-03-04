package tree_test

import (
	"go-datastruct/tree"
	"testing"
)

func validateTargets(label string, targets []int, results []int, t *testing.T) {
	if len(targets) != len(results) {
		t.Errorf("%s: expect count is %d, but got %d", label, len(targets), len(results))
	}

	for index, target := range targets {
		if target != results[index] {
			t.Errorf("%s: expect %d is %d, but got %d", label, index, target, results[index])
		}
	}
}

func TestBFS(t *testing.T) {
	st := tree.NewSearchTree()

	st.Insert(60)
	st.Insert(50)
	st.Delete(60)
	st.Insert(30)
	st.Insert(80)
	st.Insert(15)
	st.Insert(10)
	st.Insert(20)
	st.Insert(40)
	st.Insert(60)
	st.Insert(55)
	st.Insert(90)
	st.Insert(70)
	st.Insert(85)
	st.Insert(85)

	t.Logf("Origin: \r\n" + tree.FormatTree(st))

	targets := []int{50, 30, 80, 15, 40, 60, 90, 10, 20, 55, 70, 85}
	bfsResults := []int{}

	st.BFS(func(element int) {
		bfsResults = append(bfsResults, element)
	})

	validateTargets("BFS", targets, bfsResults, t)
}
