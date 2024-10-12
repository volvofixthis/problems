package first

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findSum(root *TreeNode) int {
	if root == nil {
		return 0
	}
	sL := findSum(root.Left)
	sR := findSum(root.Right)
	return root.Val + sL + sR
}

func absDiff(a, b int) int {
	if a < b {
		return b - a
	}
	return a - b
}

func findTilt(root *TreeNode) int {
	if root == nil {
		return 0
	}
	sL := findSum(root.Left)
	sR := findSum(root.Right)
	currentTilt := absDiff(sL, sR)

	return currentTilt + findTilt(root.Left) + findTilt(root.Right)
}
