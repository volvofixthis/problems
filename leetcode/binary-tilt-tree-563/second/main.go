package first

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func absDiff(a, b int) int {
	if a < b {
		return b - a
	}
	return a - b
}

func dfs(root *TreeNode, tilt *int) int {
	if root == nil {
		return 0
	}
	sL := dfs(root.Left, tilt)
	sR := dfs(root.Right, tilt)

	*tilt += absDiff(sL, sR)
	return root.Val + sL + sR
}

func findTilt(root *TreeNode) int {
	tilt := 0

	dfs(root, &tilt)

	return tilt
}
