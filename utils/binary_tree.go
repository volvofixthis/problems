package utils

import (
	"fmt"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func TreeVal(v int) *int {
	return &v
}

var a []*int = []*int{TreeVal(1), nil}

func ArrToBinaryTree(arr []*int, root *TreeNode, i, n int) *TreeNode {
	if i < n {
		if arr[i] == nil {
			return root
		}
		root = &TreeNode{Val: *arr[i]}

		root.Left = ArrToBinaryTree(arr, root.Left, i*2+1, n)
		root.Right = ArrToBinaryTree(arr, root.Right, i*2+2, n)
	}
	return root
}

func PrintBinaryTree(root *TreeNode, space int) {
	if root == nil {
		return
	}
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		cl := len(queue)
		for _, v := range queue {
			if v != nil {
				fmt.Print(v.Val)
			}
			fmt.Print(strings.Repeat(" ", space))
		}
		fmt.Println()
		for range cl {
			cur := queue[0]
			queue = queue[1:]
			if cur == nil {
				continue
			}
			queue = append(queue, cur.Left)
			queue = append(queue, cur.Right)
		}
	}
}

func PrintBinaryTree2(root *TreeNode, prefix string, isLeft bool) {
	if root != nil {
		fmt.Println(prefix + (ifThenElse(isLeft, "├── ", "└── ").(string)) + fmt.Sprint(root.Val))
		PrintBinaryTree2(root.Left, prefix+ifThenElse(isLeft, "│   ", "    ").(string), true)
		PrintBinaryTree2(root.Right, prefix+ifThenElse(isLeft, "│   ", "    ").(string), false)
	}
}

// Utility function for conditional strings (like ternary operator)
func ifThenElse(cond bool, a, b interface{}) interface{} {
	if cond {
		return a
	}
	return b
}
