package easy

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return recursive(root.Left, root.Right)
}

func recursive(left, right *TreeNode) bool {
	if left.Val != right.Val {
		return false
	}
	if left == nil && right != nil {
		return false
	}
	if left != nil && right == nil {
		return false
	}
	recursive(left.Left, right.Right)
	recursive(left.Right, right.Left)
	return true
}
