package tree

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil || root == p || root == q {
		return root
	}
	// 跟左子树要信息
	left := lowestCommonAncestor(root.Left, p, q)
	// 跟右子树要信息
	right := lowestCommonAncestor(root.Right, p, q)
	// 如果左子树不为空且又子树也不为空，则返回当前节点
	if left != nil && right != nil {
		return root
	}
	if left == nil {
		return right
	} else {
		return left
	}
}
