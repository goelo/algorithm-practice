package main

// 重点在于画图
func swapPairs(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	dummy := &ListNode{
		Next: head,
	}

	pre := dummy
	for head != nil && head.Next != nil {
		// 头节点移动
		pre.Next = head.Next
		// 第1个节点指向第三个节点
		tmp := head.Next.Next
		head.Next.Next = head
		// 第二个节点指向前一个节点
		head.Next = tmp
		// 移动pre
		pre = head
		// 移动head
		head = tmp
	}
	return dummy.Next
}
