package tree

func convertBST(root *TreeNode) *TreeNode {
	var dfs func(root *TreeNode)
	var sum int
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		dfs(root.Right)
		sum += root.Val
		root.Val = sum
		dfs(root.Left)
	}
	dfs(root)
	return root
}
