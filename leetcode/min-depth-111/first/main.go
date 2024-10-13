package main

import (
	"github.com/volvofixthis/problems/utils"
)

func minDepth(root *utils.TreeNode) int {
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

func main() {
	// arr := []*int{
	// 	utils.TreeVal(3),
	// 	utils.TreeVal(9),
	// 	utils.TreeVal(20),
	// 	// Childs of 9
	// 	nil,
	// 	nil,
	// 	// Childs of 20
	// 	utils.TreeVal(15),
	// 	utils.TreeVal(7),
	// 	nil, nil,
	// 	nil, nil,
	// 	// Childs of 15
	// 	utils.TreeVal(13),
	// 	utils.TreeVal(56),
	// 	// Childs of 7
	// 	utils.TreeVal(111),
	// 	utils.TreeVal(114)}
	// tree := utils.ArrToBinaryTree(arr, nil, 0, len(arr))
	// utils.PrintBinaryTree(tree, 4)
	// utils.PrintBinaryTree2(tree, "", false)
	arr2 := []*int{utils.TreeVal(2), nil, utils.TreeVal(3), nil, utils.TreeVal(4), nil, utils.TreeVal(5), nil, utils.TreeVal(6)}
	tree := utils.ArrToBinaryTree2(arr2)
	utils.PrintBinaryTree(tree, 4)
	utils.PrintBinaryTree2(tree, "", false)
}
