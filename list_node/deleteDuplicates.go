package main

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	pre := &ListNode{0, head}
	cur := pre

	for cur.Next != nil && cur.Next.Next != nil {
		if cur.Next.Val != cur.Next.Next.Val {
			cur = cur.Next
		} else {
			val := cur.Next.Val
			for cur.Next != nil && cur.Next.Val == val {
				// 这里很巧妙，不停的跳过相等的节点
				cur.Next = cur.Next.Next
			}
		}
	}
	return pre.Next
}
