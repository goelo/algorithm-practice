package tree

func diameterOfBinaryTree(root *TreeNode) int {
	var res int
	var dfs func(root *TreeNode) int
	dfs = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		// 获取左子树的深度,以左节点为根节点时的深度
		l := dfs(root.Left)
		// 获取右子树的深度，以右节点为根节点时的深度
		r := dfs(root.Right)
		res = max(res, l+r)
		// 此时需要返回当前节点作为根节点时，深度
		return max(l, r) + 1 // 此处的1是当前节点作为根节点时，从0开始。同root计算深度一致
	}
	dfs(root)
	return res
}
