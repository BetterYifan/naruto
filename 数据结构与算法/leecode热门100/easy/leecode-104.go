package easy

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return IfThen(maxDepth(root.Left), maxDepth(root.Right)) + 1
}

func IfThen(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
