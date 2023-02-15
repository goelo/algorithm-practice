package main

// 官方的最优解, 中点+链表逆序(双指针记录前后节点做法)+合并链表
func reorderList(head *ListNode) {
	if head == nil {
		return
	}

	slow := head
	fast := head
	// 快慢指针找到中间位置, 循环结束slow为中间位置
	// 这里需要注意判断退出的条件，fast.Next.Next对应为偶数个节点的时候走到了倒数第二个节点位置，此时slow指向中点
	// fast.Next为奇数个节点，此时走到了尾节点，slow指向中点
	// 公式性的内容记住即可。注意都是从头节点开始
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	// 截断中间位置，分成两个链表
	right := slow.Next
	slow.Next = nil
	// 定义栈存放有半部分的节点
	stack := make([]*ListNode, 0)
	for right != nil {
		node := right
		right = right.Next
		node.Next = nil
		stack = append(stack, node)
	}

	pre := head
	for pre != nil {
		// 弹出节点
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		// 临时记录当前pre的next节点
		tmp := pre.Next
		// 更改当前pre的next指向
		pre.Next = node
		// 将node的next指向当前pre的原来的next节点
		node.Next = tmp
		// pre移动到原来的next节点上
		pre = tmp
	}
}
