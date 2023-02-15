package tree

// 第一种解法 使用递归, 先前序遍历，然后调节节点即可 空间复杂度不满足o1
func flatten1(root *TreeNode) {
	// 先前序遍历，得到结果
	res := preOrder(root)
	// 这里取0...n-2的范围，只需要调节第一个至n-1个的右节点，最后一个节点无需调整
	for i := 0; i < len(res)-1; i++ {
		curr := res[i]
		next := res[i+1]
		curr.Left, curr.Right = nil, next
	}
}
func preOrder(root *TreeNode) []*TreeNode {
	var res []*TreeNode
	var preorder func(root *TreeNode)
	preorder = func(root *TreeNode) {
		if root == nil {
			return
		}
		// 根左右
		res = append(res, root)
		preorder(root.Left)
		preorder(root.Right)
	}
	preorder(root)
	return res
}

// 第二种解法，按照前序遍历的根左右先拆分
// 按照根 左子树 右子树拆分 然后将左子树和右子树分别处理成链表，然后最后链接。
// 内部左右子树又需要按照前序遍历拆分成根左右 因此自然考虑使用递归。这种属于自顶向下求解

func flatten2(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	// 先找左子树
	l := flatten2(root.Left)
	// 再找右子树
	r := flatten2(root.Right)
	if l == nil {
		root.Right = r
	} else {
		root.Right = l
		curr := l
		for curr.Right != nil {
			curr = curr.Right
		}
		curr.Right = r
		root.Left = nil
	}
	return root
}

// 第三种解法，类似倒序。自底向上，使用一个last指针记录信息，然后拼接起来
