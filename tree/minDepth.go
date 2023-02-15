package tree

import "math"

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return 1
	}
	minN := math.MaxInt64
	if root.Left != nil {
		// 拿到左子树的最小深度
		minN = min(minDepth(root.Left), minN)
	}
	if root.Right != nil {
		// 拿到右子树的最小深度
		minN = min(minDepth(root.Right), minN)
	}
	return minN + 1
}
