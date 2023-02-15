package tree

// 常规解法
func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}
	res := 1
	if root.Left != nil {
		res += countNodes(root.Left)
	}
	if root.Right != nil {
		res += countNodes(root.Right)
	}
	return res
	//var getNums func(root *TreeNode)
	//getNums = func(root *TreeNode) {
	//	if root == nil {
	//		return
	//	}
	//	if root.Left != nil {
	//		getNums(root.Left)
	//		res++
	//	}
	//	if root.Right != nil {
	//		getNums(root.Right)
	//		res++
	//	}
	//}
	//
	//getNums(root)
	//// 这里需要注意记录根节点
	//return res + 1
}

// https://programmercarl.com/0222.%E5%AE%8C%E5%85%A8%E4%BA%8C%E5%8F%89%E6%A0%91%E7%9A%84%E8%8A%82%E7%82%B9%E4%B8%AA%E6%95%B0.html#go
// 依靠完全二叉树的特点来求解
func countTreeNodes(root *TreeNode) int {
	//
	if root == nil {
		return 0
	}
	var (
		ldepth int
		rdepth int
	)
	// 注意下面记录的分别是各自的左右子树，无论在哪一层，都是各自节点作为根节点之后的左右子树。
	// 也就是判断满二叉树
	// 在完全二叉树中，如果递归向左遍历的层数等于递归向右遍历的层数，那说明就是满二叉树。
	// 计算左子树的层数
	for root.Left != nil {
		root.Left = root.Left.Left
		ldepth++
	}
	// 计算右子树的层数
	for root.Right != nil {
		root.Right = root.Right.Right
		rdepth++
	}
	// 左右子树到达各自树的最底层之后，如果左右子树层数相等, 则计算整棵子树的节点个数返回
	// 如果满二叉树的层数为h，则总节点数为：2^h - 1，根节点为第一层
	if ldepth == rdepth {
		// 这里记录的是整棵树的节点数
		return 2<<ldepth - 1 // 注意(2<<1) 相当于2^2，所以leftDepth初始为0
	}
	// 如果不相等，那么就统计左右子树的总节点,然后相加
	l := countNodes(root.Left) // 递归之后看下一层的节点作为根节点之后，是否存在满二叉树
	r := countNodes(root.Right)
	return l + r + 1
}
