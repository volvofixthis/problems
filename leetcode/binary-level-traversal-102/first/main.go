package first

import "github.com/volvofixthis/problems/utils"

func levelOrder(root *utils.TreeNode) [][]int {
	m := [][]int{}
	if root == nil {
		return m
	}
	queue := []*utils.TreeNode{root}
	for len(queue) > 0 {
		cl := len(queue)
		line := []int{}
		for _, v := range queue {
			line = append(line, v.Val)
		}
		m = append(m, line)
		for range cl {
			current := queue[0]
			queue = queue[1:]
			if current.Left == nil {
				queue = append(queue, current.Left)
			}
			if current.Right == nil {
				queue = append(queue, current.Right)
			}
		}
	}
	return m
}
