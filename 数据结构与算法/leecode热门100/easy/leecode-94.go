package easy

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//中序遍历
func inorderTraversal(root *TreeNode) []int {
	res := make([]int, 0)
	var inner func(node *TreeNode)
	inner = func(node *TreeNode) {
		if node == nil {
			return
		}
		inner(node.Left)
		res = append(res, node.Val)
		inner(node.Right)
	}
	inner(root)
	return res
}
