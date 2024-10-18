package first

import "github.com/volvofixthis/problems/utils/binarytree"

func maxDepth(root *binarytree.TreeNode) int {
	m := 0
	if root == nil {
		return m
	}
	queue := []*binarytree.TreeNode{root}
	for len(queue) > 0 {
		m++
		cl := len(queue)
		for range cl {
			current := queue[0]
			queue = queue[1:]
			if current.Left != nil {
				queue = append(queue, current.Left)
			}
			if current.Right != nil {
				queue = append(queue, current.Right)
			}
		}
	}
	return m
}
