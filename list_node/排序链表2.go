package main

func sortList2(head *ListNode) *ListNode {
	// 先找中点
	// 这里边界条件判断是当head==nil, 很容易理解。但是不能忽视head.Next != nil
	// 核心是下面截断用的是slow.Next,不判断head.Next，就会死循环。
	// 其实也很好理解，分割到只剩下一个节点的时候，无法再分，这个时候递归的退出条件就是当前节点的下一个节点是nil
	if head == nil || head.Next == nil {
		return head // 这里注意返回的是head
	}

	slow, fast := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	// 找到中点以后分割
	mid := slow.Next
	slow.Next = nil
	l := sortList2(head)
	r := sortList2(mid)
	// 归并排序
	return mergeList2(l, r)
}

func mergeList2(l, r *ListNode) *ListNode {
	dummy := &ListNode{}
	cur := dummy
	p1, p2 := l, r
	for p1 != nil && p2 != nil {
		if p1.Val < p2.Val {
			cur.Next = p1
			p1 = p1.Next
		} else {
			cur.Next = p2
			p2 = p2.Next
		}
		cur = cur.Next
	}
	if p1 == nil {
		cur.Next = p2
	}
	if p2 == nil {
		cur.Next = p1
	}
	return dummy.Next
}
