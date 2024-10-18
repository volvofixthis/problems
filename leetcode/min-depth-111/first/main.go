package main

import (
	"github.com/volvofixthis/problems/utils/binarytree"
)

func minDepth(root *binarytree.TreeNode) int {
	if root == nil {
		return 0
	}
	m := 0
	if root.Left != nil {
		m = minDepth(root.Left)
	}
	if root.Right != nil {
		lR := minDepth(root.Right)
		if m == 0 || m > lR {
			m = lR
		}
	}
	return m + 1
}
