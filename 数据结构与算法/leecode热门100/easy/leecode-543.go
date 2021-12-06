package easy

// 两个叶子节点的路径等于根节点的左右儿子深度之和,
//    2
//  3    4
//  2的深度：max(3, 4的深度)+1
func diameterOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}
	// m用来传值
	m := 0
	depth(root, &m)
	return m
}

func depth(root *TreeNode, m *int) int {
	if root == nil {
		return 0
	}
	left := depth(root.Left, m)
	right := depth(root.Right, m)
	*m = max(left+right, *m)
	return max(left, right) + 1
}

func max(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}
