package first

import "github.com/volvofixthis/problems/utils"

func levelOrder(root *utils.TreeNode) [][]int {
	m := [][]int{}
	if root == nil {
		return m
	}
	queue := []*utils.TreeNode{root}
	m = append(m, []int{root.Val})
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
		if len(line) > 0 {
			m = append(m, line)
		}
	}
	// Revert array
	for i, j := 0, len(m)-1; i < j; i, j = i+1, j-1 {
		m[i], m[j] = m[j], m[i]
	}
	return m
}
