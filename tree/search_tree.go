package tree

type SearchTreeNode struct {
	Element int
	Left    *SearchTreeNode
	Right   *SearchTreeNode
}

type SearchTree struct {
	root *SearchTreeNode
}

func NewSearchTreeNode(element int) *SearchTreeNode {
	return &SearchTreeNode{
		Element: element,
		Left:    nil,
		Right:   nil,
	}
}

func NewSearchTree() *SearchTree {
	return &SearchTree{
		root: nil,
	}
}

func (tree *SearchTree) Find(element int) *SearchTreeNode {
	node := tree.root
	for {
		if node == nil {
			return nil
		}
		if element > node.Element {
			node = node.Right
		} else if element < node.Element {
			node = node.Left
		} else {
			return node
		}
	}
}

func (tree *SearchTree) FindMin() *SearchTreeNode {
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

func (tree *SearchTree) FindMax() *SearchTreeNode {
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

func (tree *SearchTree) Insert(element int) {
	node := tree.root
	if node == nil {
		tree.root = NewSearchTreeNode(element)
		return
	}
	for {
		if element > node.Element {
			if node.Right == nil {
				node.Right = NewSearchTreeNode(element)
				return
			}
			node = node.Right
		} else if element < node.Element {
			if node.Left == nil {
				node.Left = NewSearchTreeNode(element)
				return
			}
			node = node.Left
		} else {
			return
		}
	}
}

func (tree *SearchTree) Delete(element int) {
	node := tree.root
	var parent *SearchTreeNode
	for {
		if node == nil {
			break
		}
		if element > node.Element {
			parent = node
			node = node.Right
		} else if element < node.Element {
			parent = node
			node = node.Left
		} else {
			break
		}
	}
	if node == nil {
		return
	}
	if node.Left != nil && node.Right != nil {
		parent = node
		minNode := node.Right
		for minNode.Left != nil {
			parent = minNode
			minNode = minNode.Left
		}
		node.Element = minNode.Element
		node = minNode
	}

	var child *SearchTreeNode
	if node.Left != nil {
		child = node.Left
	} else if node.Right != nil {
		child = node.Right
	}

	if parent == nil {
		tree.root = child
	} else if parent.Left == node {
		parent.Left = child
	} else if parent.Right == node {
		parent.Right = child
	}

}
