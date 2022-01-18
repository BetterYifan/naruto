package easy

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func mirrorTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	if root.Left != nil || root.Right != nil {
		root.Left, root.Right = root.Right, root.Left
	}
	mirrorTree(root.Left)
	mirrorTree(root.Right)
	return root
}
