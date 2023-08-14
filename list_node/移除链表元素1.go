package main

func removeElement(head *ListNode, val int) *ListNode {
	if head == nil {
		return nil
	}
	cur := &ListNode{
		Next: head,
	}
	// !!! 这里要记录一个头节点
	dummy := cur
	for cur.Next != nil {
		if cur.Next.Val == val {
			// 移除
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}
	}
	return dummy.Next
}
