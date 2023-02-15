package tree

func rob(root *TreeNode) int {
	var dfs func(root *TreeNode) []int

	dfs = func(root *TreeNode) []int {
		// 终止条件
		if root == nil {
			return []int{0, 0}
		}
		// 使用后序遍历, 可以得到左子树和右子树的最大值
		// 0: 偷当前节点的最大值，1: 不偷当前节点的最大值
		left := dfs(root.Left)
		right := dfs(root.Right)
		// 如果偷当前节点，那么就不考虑左右孩子.
		// 注意这个顺序，偷当前节点，那么就应该判断左右孩子节点不偷时的最大值，因此是1
		chose := root.Val + left[1] + right[1]
		// 如果不偷当前节点, 那么左右孩子可偷可不偷，因此要判断如果不偷时，左右孩子子树的最大值
		noChose := max(left[0], left[1]) + max(right[0], right[1])
		return []int{chose, noChose}
	}

	res := dfs(root)
	return max(res[0], res[1])
}
