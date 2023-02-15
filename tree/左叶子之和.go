package tree

// 更优解法
func sumOfLeftLeaves2(root *TreeNode) (res int) {
	if root == nil {
		return 0
	}

	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if root.Left != nil && root.Left.Left != nil && root.Left.Right != nil {
			res += root.Left.Val
		}
		if root.Left != nil {
			dfs(root.Left)
		}
		if root.Right != nil {
			dfs(root.Right)
		}
	}
	dfs(root)
	return res
}
func sumOfLeftLeaves(root *TreeNode) (res int) {
	if root == nil {
		return 0
	}

	var dfs func(root *TreeNode, isLeft bool)
	dfs = func(root *TreeNode, isLeft bool) {
		if root.Left == nil && root.Right == nil && isLeft {
			res += root.Val
		}
		if root.Left != nil {
			dfs(root.Left, true)
		}
		if root.Right != nil {
			dfs(root.Right, false)
		}
	}
	dfs(root, false)
	return res
}
