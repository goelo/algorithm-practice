package tree

// 反转数组,这个方法记下来,双指针法
func reverse(nums []int) []int {
	l := 0
	r := len(nums) - 1
	for l < r {
		nums[l], nums[r] = nums[r], nums[l]
		l++
		r--
	}
	return nums
}

func zigzagLevelOrder(root *TreeNode) (ans [][]int) {
	if root == nil {
		return nil
	}
	// 宽度遍历用队列
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	depth := 0
	// 根节点提前记下来
	for len(queue) > 0 {
		// 记录当前层的节点值
		curVal := make([]int, 0)
		size := len(queue)
		// 判断当前层是否处理完所有节点
		for size > 0 {
			// 弹出节点
			node := queue[0]
			queue = queue[1:]
			// 每次弹出的节点记下来，这样每一层的节点都会记录下来
			curVal = append(curVal, node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
			size--
		}
		// 此时处理[2,3]
		// 当是depth是奇数层的时候, 翻转
		// 当是depth是偶数层的时候，不翻转
		if depth%2 == 0 {
			ans = append(ans, curVal)
		} else {
			ans = append(ans, reverse(curVal))
		}
		depth++
	}
	return ans
}
