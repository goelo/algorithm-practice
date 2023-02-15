package tree

type Tree429 struct {
	Val      int
	Children []*Tree429
}

// 429 以及BFS Tree的另外一种写法
func anotherLO(root *Tree429) (ans [][]int) {
	if root == nil {
		return nil
	}
	queue := []*Tree429{root}
	for queue != nil {
		res := []int{}
		levelNodes := queue
		queue = nil
		// 遍历当前层的节点
		for _, node := range levelNodes {
			res = append(res, node.Val)
			queue = append(queue, node.Children...)
		}
		ans = append(ans, res)
	}
	return ans
}

// lc429
func levelorder(root *TreeNode) (ans [][]int) {
	if root == nil {
		return nil
	}
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	for len(queue) > 0 {
		size := len(queue)
		res := make([]int, 0, size)
		for size > 0 {
			node := queue[0]
			queue = queue[1:]
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
			res = append(res, node.Val)
			size--
		}
		ans = append(ans, res)
	}
	return ans
}

// lc102
func levelOrder(root *TreeNode) (ans [][]int) {
	if root == nil {
		return nil
	}
	// 宽度遍历，用队列
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	for len(queue) != 0 {
		size := len(queue)
		combo := make([]int, size)
		// 遍历当前层元素
		for size > 0 {
			// 弹出队列元素
			node := queue[0]
			queue = queue[1:]
			// 先左后右遍历
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
			combo = append(combo, node.Val)
			size--
		}
		ans = append(ans, combo)
	}
	return ans
}
