package main

// lc61
func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || k == 0 || head.Next == nil {
		return head
	}
	cur := head
	// 链表先结成环状,并统计链表长度，为了后续计算k位的移动
	n := 1
	for cur.Next != nil {
		cur = cur.Next
		n++
	}
	// 如果k能整除链表长度, 说明移动完之后还是原来位置
	if k%n == 0 {
		return head
	}
	// 组成环形链表，此时cur指向尾部节点
	cur.Next = head
	// 不能整除则需要确定头节点的位置，以及尾部节点移动到何处
	// 需要找到移动k位置之后新的尾部节点是哪个。按照原来的单链表来计算，即尾部节点移动n-k位置后指向的节点
	// 当组成环形之后，移动n-k个位置之后的next即是新链表的头节点。
	// !!! 这里当初看完答案做题的时候还是没有思考正确。是尾部节点需要移动n-k个位置才能找到。
	for i := 0; i < n-k%n; i++ {
		cur = cur.Next
	}
	// 此时cur移动到的节点即是尾部节点，将其next指向nil,截断链表
	dummy := &ListNode{}
	dummy.Next = cur.Next
	cur.Next = nil

	return dummy.Next
}
