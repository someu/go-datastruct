package tree_test

import (
	"go-datastruct/tree"
	"testing"
)

func TestAVLTree(t *testing.T) {
	st := tree.NewAVLTree()

	if st.FindMin() != nil {
		t.Errorf("search tree's min expected to be nil, but got %d", st.FindMin().Element)
	}

	if st.FindMax() != nil {
		t.Errorf("search tree's max expected to be nil, but got %d", st.FindMax().Element)
	}

	st.Insert(55)
	t.Logf("Origin: \r\n" + tree.FormatAVLTree(st))
	st.Insert(60)
	t.Logf("Origin: \r\n" + tree.FormatAVLTree(st))
	st.Insert(57)
	t.Logf("Origin: \r\n" + tree.FormatAVLTree(st))
	st.Insert(70)
	t.Logf("Origin: \r\n" + tree.FormatAVLTree(st))
	st.Insert(80)
	t.Logf("Origin: \r\n" + tree.FormatAVLTree(st))
	st.Insert(85)
	t.Logf("Origin: \r\n" + tree.FormatAVLTree(st))
	st.Insert(90)
	t.Logf("Origin: \r\n" + tree.FormatAVLTree(st))
	st.Insert(10)
	t.Logf("Origin: \r\n" + tree.FormatAVLTree(st))
	st.Insert(15)
	t.Logf("Origin: \r\n" + tree.FormatAVLTree(st))
	st.Insert(20)
	t.Logf("Origin: \r\n" + tree.FormatAVLTree(st))
	st.Insert(30)
	t.Logf("Origin: \r\n" + tree.FormatAVLTree(st))
	st.Insert(40)
	t.Logf("Origin: \r\n" + tree.FormatAVLTree(st))
	st.Insert(50)
	t.Logf("Origin: \r\n" + tree.FormatAVLTree(st))
	st.Insert(87)
	t.Logf("Origin: \r\n" + tree.FormatAVLTree(st))

	if st.FindMax().Element != 90 {
		t.Errorf("search tree's max expected to be 90, but got %d", st.FindMax().Element)
	}
	if st.FindMin().Element != 10 {
		t.Errorf("search tree's min expected to be 10, but got %d", st.FindMin().Element)
	}

	t.Logf("Origin: \r\n" + tree.FormatAVLTree(st))

	st.Delete(10)
	node10 := st.Find(10)
	if node10 != nil {
		t.Errorf("node 10 expected to be nil, but got %d", node10.Element)
	}
	node90 := st.Find(90)
	if node90 == nil {
		t.Errorf("node 90 expected to be 90, but got nil")
	} else if node90.Element != 90 {
		t.Errorf("node 90 expected to be 90, but got %d", node90.Element)
	}
	st.Delete(90)

	t.Logf("Delete 10\r\n" + tree.FormatAVLTree(st))
	st.Delete(90)
	t.Logf("Delete 90\r\n" + tree.FormatAVLTree(st))
	st.Delete(80)
	t.Logf("Delete 80\r\n" + tree.FormatAVLTree(st))
	st.Delete(50)
	t.Logf("Delete 50\r\n" + tree.FormatAVLTree(st))
	st.Delete(57)
	t.Logf("Delete 57\r\n" + tree.FormatAVLTree(st))
	st.Delete(55)
	t.Logf("Delete 55\r\n" + tree.FormatAVLTree(st))
	st.Delete(60)
	t.Logf("Delete 60\r\n" + tree.FormatAVLTree(st))
}
