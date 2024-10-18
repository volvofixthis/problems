package first

import "github.com/volvofixthis/problems/utils/binarytree"

func zigzagLevelOrder(root *binarytree.TreeNode) [][]int {
	m := [][]int{}
	if root == nil {
		return m
	}
	queue := []*binarytree.TreeNode{root}
	reversed := false
	for len(queue) > 0 {
		cl := len(queue)
		line := make([]int, 0, cl)
		for range cl {
			current := queue[0]
			queue = queue[1:]
			if current.Left != nil {
				queue = append(queue, current.Left)
				line = append(line, current.Left.Val)
			}
			if current.Right != nil {
				queue = append(queue, current.Right)
				line = append(line, current.Right.Val)
			}
		}
		if reversed {
			for i, j := 0, len(line)-1; i < j; i, j = i+1, j-1 {
				line[i], line[j] = line[j], line[i]
			}
		}
		if len(line) > 0 {
			m = append(m, line)
		}
		reversed = !reversed
	}
	return m
}
