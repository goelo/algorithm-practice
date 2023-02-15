package tree

func levelOrderBottom(root *TreeNode) (res [][]int) {
	if root == nil {
		return nil
	}
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	for len(queue) > 0 {
		size := len(queue)
		// 记录当前层的节点值
		ans := make([]int, size)
		for size > 0 {
			node := queue[0]
			queue = queue[1:]
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
			ans = append(ans, node.Val)
			size--
		}
		res = append(res, ans)
	}
	Reverse(res)
	return res
}

func Reverse(nums [][]int) [][]int {
	i := 0
	j := len(nums) - 1
	for i < j {
		nums[i], nums[j] = nums[j], nums[i]
		i++
		j--
	}
	return nums
}
