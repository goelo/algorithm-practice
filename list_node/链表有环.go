package main

func detectCycle1(head *ListNode) *ListNode {
	fast, slow := head, head
	for fast != nil && fast.Next != nil {
		// fast每次走两步, slow每次走一步
		fast = fast.Next.Next
		slow = slow.Next
		// 如果相遇
		if fast == slow {
			// slow回到头
			slow = head
			// 如果不能相遇
			for slow != fast {
				slow = slow.Next
				fast = fast.Next
			}
			return fast
		}
	}
	return nil
}
