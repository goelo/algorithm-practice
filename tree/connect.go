package tree

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	if root == nil {
		return nil
	}
	queue := make([]*Node, 0)
	queue = append(queue, root)
	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			// 弹出队头元素, 然后队列缩小
			node := queue[0]
			queue = queue[1:]
			// 如果队列还有数据, 说明当前层未到末尾
			// 这里的逻辑要记住。判断当前是否是最后一个节点,如果是最后一个节点，那么next指针则设置未空。
			if i != size-1 {
				node.Next = queue[0]
			} else {
				node.Next = nil
			}
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
	}
	return root
}
