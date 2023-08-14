package main

func removeNthFromEnd1(head *ListNode, n int) *ListNode {
	dummy := &ListNode{
		Next: head,
	}
	fast, slow := dummy, dummy
	// 这种做法先移动fast n步
	for i := n; i > 0; i-- {
		fast = fast.Next
	}
	// 这里是判断fast走到最后一个节点，不能继续往下走。所以判断next是否为空
	for fast.Next != nil {
		// 同时开始走
		fast = fast.Next
		slow = slow.Next
	}
	slow.Next = slow.Next.Next
	return dummy.Next
}
