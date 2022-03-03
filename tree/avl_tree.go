package tree

type AVLTreeNode struct {
	Element int
	Height  int
	Left    *AVLTreeNode
	Right   *AVLTreeNode
}

type AVLTree struct {
	root *AVLTreeNode
}

func NewAVLTreeNode(element int) *AVLTreeNode {
	return &AVLTreeNode{
		Element: element,
		Height:  0,
		Left:    nil,
		Right:   nil,
	}
}

func NewAVLTree() *AVLTree {
	return &AVLTree{
		root: nil,
	}
}

func (tree *AVLTree) Find(element int) *AVLTreeNode {
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

func (node *AVLTreeNode) height() int {
	if node == nil {
		return -1
	}
	return node.Height
}

func (node *AVLTreeNode) fixHeight() {
	right := -1
	left := -1
	if node.Left != nil {
		left = node.Left.Height
	}
	if node.Right != nil {
		right = node.Right.Height
	}
	if right > left {
		node.Height = right + 1
	} else {
		node.Height = left + 1
	}
}

/**
	a
   /
  B      ---->        B
 /                   / \
c                   c   a
*/
func rRotation(node *AVLTreeNode) *AVLTreeNode {
	child := node.Left
	node.Left = child.Right
	child.Right = node

	node.fixHeight()
	child.fixHeight()
	return child
}

/**
a
 \
  B      ---->        B
   \                 / \
    c               a   c
*/
func lRotation(node *AVLTreeNode) *AVLTreeNode {
	child := node.Right
	node.Right = child.Left
	child.Left = node

	node.fixHeight()
	child.fixHeight()
	return child
}

/**
a
 \
  B      ---->        c
 /                   / \
c                   a   B
*/
func rlRotation(node *AVLTreeNode) *AVLTreeNode {
	node.Right = rRotation(node.Right)
	return lRotation(node)
}

/**
  a
 /
B      ---->       c
 \                / \
  c              B   a
*/
func lrRotation(node *AVLTreeNode) *AVLTreeNode {
	node.Left = lRotation(node.Left)
	return rRotation(node)
}

func handleBF(node *AVLTreeNode) *AVLTreeNode {
	if node.Left.height()-node.Right.height() == 2 {
		if node.Left.Left.height() > node.Left.Right.height() {
			// ll
			return rRotation(node)
		} else {
			// lr
			return lrRotation(node)
		}
	} else if node.Left.height()-node.Right.height() == -2 {
		if node.Right.Left.height() > node.Right.Right.height() {
			// rl
			return rlRotation(node)
		} else {
			// rr
			return lRotation(node)
		}
	} else {
		node.fixHeight()
		return node
	}
}

func insert(node *AVLTreeNode, element int) *AVLTreeNode {
	if node == nil {
		return NewAVLTreeNode(element)
	}
	if element > node.Element {
		node.Right = insert(node.Right, element)
	} else if element < node.Element {
		node.Left = insert(node.Left, element)
	}
	return handleBF(node)
}

func (tree *AVLTree) Insert(element int) {
	tree.root = insert(tree.root, element)
}

func delete(node *AVLTreeNode, element int) *AVLTreeNode {
	if node == nil {
		return nil
	}
	if node.Element < element {
		node.Right = delete(node.Right, element)
	} else if node.Element > element {
		node.Left = delete(node.Left, element)
	} else {
		if node.Left == nil && node.Right == nil {
			return nil
		} else if node.Left == nil && node.Right != nil {
			return node.Right
		} else if node.Left != nil && node.Right == nil {
			return node.Left
		} else {
			n := node.Left
			for n.Right != nil {
				n = n.Right
			}
			node.Element = n.Element
			node.Left = delete(node.Left, element)
		}
	}
	return handleBF(node)
}

func (tree *AVLTree) Delete(element int) {
	tree.root = delete(tree.root, element)
}
