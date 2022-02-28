package tree

import "math"

type AVLTreeNode struct {
	Element interface{}
	Height  int
	Left    *AVLTreeNode
	Right   *AVLTreeNode
}

type AVLTree struct {
	root    *AVLTreeNode
	compare Compare
}

func NewAVLTreeNode(element interface{}) *AVLTreeNode {
	return &AVLTreeNode{
		Element: element,
		Height:  0,
		Left:    nil,
		Right:   nil,
	}
}

func NewAVLTree(compare Compare) *AVLTree {
	return &AVLTree{
		root:    nil,
		compare: compare,
	}
}

func (tree *AVLTree) Find(element interface{}) *AVLTreeNode {
	node := tree.root
	for {
		if node == nil {
			return nil
		}
		compared := tree.compare(element, node.Element)
		if compared > 0 {
			node = node.Right
		} else if compared < 0 {
			node = node.Left
		} else {
			return node
		}
	}
}

func (tree *AVLTree) FindMin() *AVLTreeNode {
	if tree.root == nil {
		return nil
	}
	node := tree.root
	for {
		if node.Left == nil {
			return node
		}
		node = node.Left
	}
}

func (tree *AVLTree) FindMax() *AVLTreeNode {
	if tree.root == nil {
		return nil
	}
	node := tree.root
	for {
		if node.Right == nil {
			return node
		}
		node = node.Right
	}
}

/**
	A
   /
  B      ---->        B
 /                   / \
C                   C   A
*/
func llRotation(node *AVLTreeNode) *AVLTreeNode {
	child := node.Left
	node.Left = child.Right
	child.Right = node

	node.Height = int(math.Max(float64(node.Left.Height), float64(node.Right.Height)))
	child.Height = int(math.Max(float64(child.Left.Height), float64(child.Right.Height)))

	return child
}

/**
	A
	 \
      B      ---->        B
       \                 / \
        C               A   C
*/
func rrRotation(node *AVLTreeNode) *AVLTreeNode {
	child := node.Right
	node.Right = child.Left
	child.Left = node

	node.Height = int(math.Max(float64(node.Left.Height), float64(node.Right.Height)))
	child.Height = int(math.Max(float64(child.Left.Height), float64(child.Right.Height)))

	return child
}

/**
	A
	 \
      B      ---->        C
     /                   / \
    C                   A   B
*/
func rlRotation(node *AVLTreeNode) *AVLTreeNode {
	child := node.Right
	node.Right = nil
	child.Left = node
	return child
}

func (node *AVLTreeNode) Insert(element interface{}) {

}

func (tree *AVLTree) Insert(element interface{}) {
	node := tree.root
	if node == nil {
		tree.root = NewAVLTreeNode(element)
		return
	}
	for {
		compared := tree.compare(element, node.Element)
		if compared > 0 {
			if node.Right == nil {
				node.Right = NewAVLTreeNode(element)
				break
			}
			node = node.Right
		} else if compared < 0 {
			if node.Left == nil {
				node.Left = NewAVLTreeNode(element)
				break
			}
			node = node.Left
		} else {
			break
		}
	}

}

func (tree *AVLTree) Delete(element interface{}) {

}

/*
		A
	   / \
	  B   D
     /
	C

	  B
	 / \
	C   A

		A
	   /
	  B
       \
	    C



*/
