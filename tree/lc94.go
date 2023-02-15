package tree

type lc94TreeNode struct {
	Val   int
	Left  *lc94TreeNode
	Right *lc94TreeNode
}

// 中序遍历 左根右
func inorderTraversal(root *lc94TreeNode) (res []int) {
	var inorder func(node *lc94TreeNode)
	inorder = func(node *lc94TreeNode) {
		if node == nil {
			return
		}
		inorder(node.Left)
		res = append(res, node.Val)
		inorder(node.Right)
	}
	inorder(root)
	return
}

// 迭代
func inorder(root *lc94TreeNode) (res []int) {
	var stack []*lc94TreeNode
	for root != nil || len(stack) != 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, root.Val)
		root = root.Right
	}
	return res
}

// 前序遍历 根左右
func preorderTraversal(root *lc94TreeNode) (res []int) {
	var preorder func(node *lc94TreeNode)
	preorder = func(node *lc94TreeNode) {
		if node == nil {
			return
		}
		res = append(res, node.Val)
		preorder(node.Left)
		preorder(node.Right)
	}
	preorder(root)
	return
}

func preorder(root *lc94TreeNode) (res []int) {
	var stack []*lc94TreeNode
	for root != nil || len(stack) != 0 {
		for root != nil {
			// 根节点值先保存, 然后遍历左子树
			res = append(res, root.Val)
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		root = root.Right
	}
	return res
}

// 后序遍历 左右根
func postorderTraversal(root *lc94TreeNode) (res []int) {
	var postorder func(node *lc94TreeNode)
	postorder = func(node *lc94TreeNode) {
		if node == nil {
			return
		}
		postorder(node.Left)
		postorder(node.Right)
		res = append(res, node.Val)
	}
	postorder(root)
	return
}

// 迭代
func postorder(root *lc94TreeNode) (res []int) {
	var stack []*lc94TreeNode
	var prev *lc94TreeNode // 上一次访问的节点
	for root != nil || len(stack) != 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}

		// 出栈
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		// 右子树为空
		// 到达叶子节点，或者已经访问过了则不需要在访问了，打印
		if root.Right == nil || prev == root.Right {
			res = append(res, root.Val)
			prev = root
			root = nil
		} else {
			stack = append(stack, root)
			root = root.Right
		}
	}
	return
}
