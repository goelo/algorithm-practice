package tree

import "math"

func largestValues(root *TreeNode) (ans []int) {
	if root == nil {
		return nil
	}

	queue := []*TreeNode{root}
	for queue != nil {
		levelNodes := queue
		queue = nil
		// 每一层记录一个最大值
		maxNum := math.MinInt64
		for _, node := range levelNodes {
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
			maxNum = max(maxNum, node.Val)
		}
		ans = append(ans, maxNum)
	}
	return ans
}
