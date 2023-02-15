package tree

import "math"

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	// 如果左子树或者右子树不是平衡二叉树，那么返回false
	if !isBalanced(root.Left) || !isBalanced(root.Right) {
		return false
	}
	var dfs func(root *TreeNode) int

	dfs = func(root *TreeNode) int {
		// 到达底部了
		if root == nil {
			return 0
		}
		l := dfs(root.Left)
		r := dfs(root.Right)
		return max(l, r) + 1
	}

	// 判断左子树高度
	l := dfs(root.Left)
	// 判断右子树高度
	r := dfs(root.Right)
	// 判断根节点是否是平衡二叉树
	if math.Abs(float64(l-r)) <= 1 {
		return true
	}
	return false
}
