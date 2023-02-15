package main

func detectCycle(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	// 这里的起始位置放在相同节点
	slow := head
	fast := head
	// 只要for循环不退出，表示链表有环
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			fast = head
			for slow != fast {
				slow = slow.Next
				fast = fast.Next
			}
			return fast
		}
	}
	return nil
}
