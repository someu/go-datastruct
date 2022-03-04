package tree

import (
	"go-datastruct/stack"
)

func recursionPreOrder(node *SearchTreeNode, fn func(int)) {
	if node == nil {
		return
	}
	fn(node.Element)
	recursionPreOrder(node.Left, fn)
	recursionPreOrder(node.Right, fn)
}

func recursionMiddleOrder(node *SearchTreeNode, fn func(int)) {
	if node == nil {
		return
	}
	recursionMiddleOrder(node.Left, fn)
	fn(node.Element)
	recursionMiddleOrder(node.Right, fn)
}

func recursionPostOrder(node *SearchTreeNode, fn func(int)) {
	if node == nil {
		return
	}
	recursionPostOrder(node.Left, fn)
	recursionPostOrder(node.Right, fn)
	fn(node.Element)
}

// 递归先序遍历
func (tree *SearchTree) DFSRecursionPreOrder(fn func(int)) {
	recursionPreOrder(tree.root, fn)
}

// 递归中序遍历
func (tree *SearchTree) DFSRecursionMiddleOrder(fn func(int)) {
	recursionMiddleOrder(tree.root, fn)
}

// 递归后序遍历
func (tree *SearchTree) DFSRecursionPostOrder(fn func(int)) {
	recursionPostOrder(tree.root, fn)
}

// 非递归先序遍历
func (tree *SearchTree) DFSPreOrder(fn func(int)) {
	s := stack.New()
	node := tree.root
	if node != nil {
		s.Push(node)
	}
	for s.Size > 0 {
		node = s.Pop().(*SearchTreeNode)
		fn(node.Element)
		if node.Right != nil {
			s.Push(node.Right)
		}
		if node.Left != nil {
			s.Push(node.Left)
		}
	}
}

// 非递归中序遍历
func (tree *SearchTree) DFSMiddleOrder(fn func(int)) {
	s := stack.New()
	node := tree.root
	for node != nil || s.Size > 0 {
		for node != nil {
			s.Push(node)
			node = node.Left
		}
		current := s.Pop().(*SearchTreeNode)
		fn(current.Element)
		node = current.Right
	}
}

// 非递归后序遍历
func (tree *SearchTree) DFSPostOrder(fn func(int)) {
	s1 := stack.New()
	s2 := stack.New()
	if tree.root != nil {
		s1.Push(tree.root)
	}
	for s1.Size > 0 {
		node := s1.Pop().(*SearchTreeNode)
		s2.Push(node)
		if node.Left != nil {
			s1.Push(node.Left)
		}
		if node.Right != nil {
			s1.Push(node.Right)
		}
	}
	for s2.Size > 0 {
		fn(s2.Pop().(*SearchTreeNode).Element)
	}
}
