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

// ArrToBinaryTree is used to load implicit representation from array
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

// ArrToBinaryTree is used to load level ordered represenation from array
func ArrToBinaryTree2(arr []*int) *TreeNode {
	if len(arr) == 0 || arr[0] == nil {
		return nil
	}

	root := &TreeNode{Val: *arr[0]}
	queue := []*TreeNode{}
	queue = append(queue, root)

	pos := 1
	for len(queue) > 0 {
		cl := len(queue)
		for range cl {
			current := queue[0]
			queue = queue[1:]
			if pos < len(arr) && arr[pos] != nil {
				n := TreeNode{
					Val: *arr[pos],
				}
				current.Left = &n
				queue = append(queue, current.Left)
			}
			pos++
			if pos < len(arr) && arr[pos] != nil {
				n := TreeNode{
					Val: *arr[pos],
				}
				current.Right = &n
				queue = append(queue, current.Right)
			}
			pos++
		}
	}

	return root
}

// PrintBinaryTree can be used to print binary tree using bfs
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

// PrintBinaryTree2 can be used to print binary tree using dfs
func PrintBinaryTree2(root *TreeNode, prefix string, isLeft bool) {
	if root != nil {
		fmt.Println(prefix + (IfThenElse(isLeft, "├── ", "└── ").(string)) + fmt.Sprint(root.Val))
		PrintBinaryTree2(root.Left, prefix+IfThenElse(isLeft, "│   ", "    ").(string), true)
		PrintBinaryTree2(root.Right, prefix+IfThenElse(isLeft, "│   ", "    ").(string), false)
	}
}

// Utility function for conditional strings (like ternary operator)
func IfThenElse(cond bool, a, b interface{}) interface{} {
	if cond {
		return a
	}
	return b
}
