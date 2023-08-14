package main

func reverseListNode(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	var pre *ListNode
	cur := head
	for cur != nil {
		tmp := cur.Next
		cur.Next = pre
		pre = cur
		cur = tmp
	}
	return pre
}
