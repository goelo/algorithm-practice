package main

// 空间复杂度是o(1)
// 双指针，但不是快慢指针，是定义两个临时指针。
// 因为反转链表，因此从头结点开始的时候，需要一个指针cur保存下一个节点的值，由于下一个节点要指向前一个结点
// 因此还需要一个指针保存前一个结点的值，来改变指向。
func RevereList(head *ListNode) *ListNode {
	var pre *ListNode // 用来保存前一个结点
	cur := head       // 移动cur结点，移动过程中记录前一个结点，同时改变后一个结点的指向
	// 一直走到链表尾部
	// 这样等走到链表尾部的时候，pre已经移动到了原来链表的尾部，链表也完成了反转
	for cur != nil {
		// 每次移动，临时定义一个临时结点来保存当前结点的下一个结点
		tmp := cur.Next
		// 截断当前cur结点的next指向，使其指向前一个结点
		cur.Next = pre
		// pre记录上一个结点
		pre = cur
		// 移动cur往下一个结点去
		cur = tmp
	}
	// 最后返回pre即可
	return pre
}

// 这种空间复杂度是o(n)
func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	// 先头节点入栈
	stack := []*ListNode{head}
	cur := head
	for cur.Next != nil {
		// 然后不停的压栈
		stack = append(stack, cur.Next)
		cur = cur.Next
	}
	dummy := &ListNode{}
	cur = dummy
	for len(stack) > 0 {
		cur.Next = stack[len(stack)-1]
		// 这里需要注意要断开每个弹出节点的next指针
		cur.Next.Next = nil
		cur = cur.Next
		stack = stack[:len(stack)-1]
	}
	return dummy.Next
}
