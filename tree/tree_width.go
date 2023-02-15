package tree

func max(a, b int) int {
	if a < b {
		return b
	} else {
		return a
	}
}

func maxWidthTree(root *TreeNode) int {
	// queue 来记录节点，queue的size则是当前层的节点个数
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	// queueIdx 记录节点的索引i, 下一层节点的索引则分别为2i， 2i+1
	queueIdx := make([]int, 0)
	queueIdx = append(queueIdx, 1)
	var (
		// 每层的第一个节点索引
		initIndex int
		// 每层的最后一个索引
		idx int
		// 最大宽度
		res int
	)
	for len(queue) != 0 {
		// 记录当前层的节点个数
		size := len(queue)
		// 记录第一个
		initIndex = queueIdx[0]
		// 2, 3
		for size > 0 {
			// 如果队列不为空，则弹出一个元素
			item := queue[0]
			queue = queue[1:]
			idx = queueIdx[0]
			queueIdx = queueIdx[1:]
			// 左子树不为空，加左子树
			if item.Left != nil {
				queue = append(queue, item.Left)
				queueIdx = append(queueIdx, 2*idx)
			}
			if item.Right != nil {
				queue = append(queue, item.Right)
				queueIdx = append(queueIdx, 2*idx+1)
			}
			size--
		}
		// 记录max
		res = max(res, idx-initIndex+1)
	}
	return res
}

// 这个是不考虑中间空节点，只是用于统计每层都有几个有效节点的宽度。
func widthOfBinaryTree(root *TreeNode) int {
	// 当前层结束节点
	// 下一层结束节点
	// 队列
	var (
		curLevelEnd  *TreeNode
		nextLevelEnd *TreeNode
		res          int
		maxN         int
	)
	q := make([]*TreeNode, 0)
	// 头节点入队
	q = append(q, root)
	// 如果队列不为空，出队列, 然后左子树入队，右子树入队
	if len(q) != 0 {
		curNode := q[0]
		q = q[1:]
		if curNode.Left != nil {
			q = append(q, curNode.Left)
			nextLevelEnd = curNode.Left
		}
		if curNode.Right != nil {
			q = append(q, curNode.Right)
			nextLevelEnd = curNode.Right
		}
		// 如果走到当前节点的最后一个
		if curNode == curLevelEnd {
			maxN = max(maxN, res+1)
			curLevelEnd = nextLevelEnd
			nextLevelEnd = nil
			res = 0
		} else {
			res += 1
		}
	}
	return maxN
}
