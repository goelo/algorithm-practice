package main

func removeLinkNode(head *ListNode, val int) *ListNode {
	if head == nil {
		return nil
	}
	dummy := &ListNode{}
	dummy.Next = head
	cur := dummy
	for dummy.Next != nil {
		if dummy.Next.Val == val {
			// 移除当前元素
			dummy.Next = dummy.Next.Next
			continue
		}
		// 当前元素后移
		dummy = dummy.Next
	}
	return cur.Next
}
