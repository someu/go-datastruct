package tree_test

import (
	"go-datastruct/tree"
	"testing"
)

func TestSearchTree(t *testing.T) {
	st := tree.NewSearchTree()

	if st.FindMin() != nil {
		t.Errorf("search tree's min expected to be nil, but got %d", st.FindMin().Element)
	}

	if st.FindMax() != nil {
		t.Errorf("search tree's max expected to be nil, but got %d", st.FindMax().Element)
	}

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

	if st.FindMax().Element != 90 {
		t.Errorf("search tree's max expected to be 90, but got %d", st.FindMax().Element)
	}
	if st.FindMin().Element != 10 {
		t.Errorf("search tree's min expected to be 10, but got %d", st.FindMin().Element)
	}

	t.Logf("Origin: \r\n" + tree.FormatTree(st))

	st.Delete(10)
	node10 := st.Find(10)
	if node10 != nil {
		t.Errorf("node 10 expected to be nil, but got %d", node10.Element)
	}
	node90 := st.Find(90)
	if node90 == nil {
		t.Errorf("node 90 expected to be 90, but got nil")
	} else if node90.Element != 90 {
		t.Errorf("node 90 expected to be nil, but got %d", node90.Element)
	}

	t.Logf("Delete 10\r\n" + tree.FormatTree(st))
	st.Delete(90)
	t.Logf("Delete 90\r\n" + tree.FormatTree(st))
	st.Delete(80)
	t.Logf("Delete 80\r\n" + tree.FormatTree(st))
	st.Delete(50)
	t.Logf("Delete 50\r\n" + tree.FormatTree(st))
	st.Delete(60)
	t.Logf("Delete 60\r\n" + tree.FormatTree(st))
	st.Delete(150)
	t.Logf("Delete 150\r\n" + tree.FormatTree(st))
}
