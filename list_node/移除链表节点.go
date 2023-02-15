package main

func removeElements(head *ListNode, val int) *ListNode {
	if head == nil {
		return head
	}

	cur := &ListNode{Next: head}
	// 遍历链表
	for cur.Next != nil {
		if cur.Next.Val == val {
			// cur结点后移
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}

	}
	return head
}
