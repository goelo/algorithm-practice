package tree

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}

	return frontOrder(root.Left, root.Right)
}

func frontOrder(l, r *TreeNode) bool {
	if l == nil && r == nil {
		return true
	}
	if l == nil || r == nil {
		return false
	}
	left, right := false, false
	// 左子树不为空，且值相等，递归
	if l.Val == r.Val {
		left = frontOrder(l.Left, r.Right)
		right = frontOrder(l.Right, r.Left)
	}

	if left && right {
		return true
	}

	return false
}
