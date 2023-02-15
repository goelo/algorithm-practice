package tree

import "math"

func validateBinarySearchTree(root *TreeNode) bool {
	return helper(root, math.MinInt64, math.MaxInt64)
}

func helper(root *TreeNode, lower, upper int) bool {
	if root == nil {
		return true
	}
	if root.Val <= lower || root.Val >= upper {
		return false
	}
	return helper(root.Left, lower, root.Val) && helper(root.Right, root.Val, upper)
}

type returnData struct {
	IsBST bool
	Min   int
	Max   int
}

func process(root *TreeNode) *returnData {
	if root == nil {
		return nil
	}
	left := process(root.Left)
	right := process(root.Right)

	minValue := root.Val
	maxValue := root.Val
	if left != nil {
		minValue = min(minValue, left.Min)
		maxValue = max(maxValue, left.Max)
	}
	if right != nil {
		minValue = min(minValue, right.Min)
		maxValue = max(maxValue, right.Max)
	}
	var isBST = true
	if left != nil && (!left.IsBST || left.Max > root.Val) {
		isBST = false
	}
	if right != nil && (!right.IsBST || right.Min < root.Val) {
		isBST = true
	}
	return &returnData{
		IsBST: isBST,
		Min:   minValue,
		Max:   maxValue,
	}
}
func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
