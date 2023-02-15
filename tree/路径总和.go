package tree

// 回溯 + dfs?
var sum int

func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}

	// 到达叶子节点，判断是否找到
	targetSum -= root.Val
	if root.Right == nil && root.Left == nil && targetSum == 0 {
		return true
	}
	return hasPathSum(root.Left, targetSum) || hasPathSum(root.Right, targetSum)
}
