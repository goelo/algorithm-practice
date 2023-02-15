package tree

func rightSideView(root *TreeNode) (res []int) {
	if root == nil {
		return nil
	}
	stack := make([]*TreeNode, 0)
	stack = append(stack, root)
	for stack != nil {
		size := len(stack)
		for i := 0; i < size; i++ {
			node := stack[0]
			stack = stack[1:]
			if node.Right != nil {
				stack = append(stack, node.Right)
			}
			if node.Left != nil {
				stack = append(stack, node.Left)
			}
			// 这里！！！判断当前层的最后一个节点, 才append
			if i == size-1 {
				res = append(res, node.Val)
			}
		}
	}
	return res
	//var dfs func(root *TreeNode)
	//dfs = func(root *TreeNode) {
	//	// 终止条件
	//	if root == nil {
	//		return
	//	}
	//	var chose bool
	//	// 右节点不为空，则把右节点加入数组
	//	if root.Right != nil {
	//		res = append(res, root.Right.Val)
	//		dfs(root.Right)
	//		chose = true
	//	}
	//	if !chose && root.Left != nil {
	//		res = append(res, root.Left.Val)
	//		dfs(root.Left)
	//	}
	//}

	//dfs(root)
	//return res
}
