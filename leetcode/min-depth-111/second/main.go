package main

import "github.com/volvofixthis/problems/utils/binarytree"

func minDepth(root *binarytree.TreeNode) int {
	if root == nil {
		return 0
	}
	queue := []*binarytree.TreeNode{}
	queue = append(queue, root)
	m := 0
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
				continue
			}
			if current.Left == nil {
				return m
			}
		}
	}
	return m
}
