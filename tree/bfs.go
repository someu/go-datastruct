package tree

import "go-datastruct/queue"

func (tree *SearchTree) BFS(fn func(int)) {
	q := queue.New()
	if tree.root != nil {
		q.Enqueue(tree.root)
	}

	for q.Size > 0 {
		node := q.Dequeue().(*SearchTreeNode)
		if node.Left != nil {
			q.Enqueue(node.Left)
		}
		if node.Right != nil {
			q.Enqueue(node.Right)
		}
		fn(node.Element)
	}
}
