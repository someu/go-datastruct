package tree

type SearchTreeNode struct {
	Element interface{}
	Left    *SearchTreeNode
	Right   *SearchTreeNode
}

type Compare func(interface{}, interface{}) int

type SearchTree struct {
	root    *SearchTreeNode
	compare Compare
}

func NewSearchTreeNode(element interface{}) *SearchTreeNode {
	return &SearchTreeNode{
		Element: element,
		Left:    nil,
		Right:   nil,
	}
}

func NewSearchTree(compare Compare) *SearchTree {
	return &SearchTree{
		root:    nil,
		compare: compare,
	}
}

func (tree *SearchTree) Find(element interface{}) *SearchTreeNode {
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

func (tree *SearchTree) Insert(element interface{}) {
	node := tree.root
	if node == nil {
		tree.root = NewSearchTreeNode(element)
		return
	}
	for {
		compared := tree.compare(element, node.Element)
		if compared > 0 {
			if node.Right == nil {
				node.Right = NewSearchTreeNode(element)
				return
			}
			node = node.Right
		} else if compared < 0 {
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

func (tree *SearchTree) Delete(element interface{}) {
	node := tree.root
	var parent *SearchTreeNode
	for {
		if node == nil {
			break
		}
		compared := tree.compare(element, node.Element)
		if compared > 0 {
			parent = node
			node = node.Right
		} else if compared < 0 {
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
		// 从右子树中取出最小值
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
