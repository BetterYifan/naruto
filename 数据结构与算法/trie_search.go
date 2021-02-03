package 数据结构与算法

/**
判断是否是二叉搜索树，利用树的中序遍历结果是有序
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}
	order := midTraverse(root)
	for i := 1; i < len(order); i++ {
		if order[i] <= order[i-1] {
			return false
		}
	}
	return true
}

func midTraverse(root *TreeNode) []int {
	var data []int
	traverse(root, &data)
	return data
}

func traverse(root *TreeNode, data *[]int) {
	if root == nil {
		return
	}
	traverse(root.Left, data)
	*data = append(*data, root.Val)
	traverse(root.Right, data)
}
