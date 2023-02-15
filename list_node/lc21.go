package main

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	head := &ListNode{}
	// 创建临时节点
	p := head // 这里思考一下, head和p都指向了最开始的头节点，所以移动p并不会影响头节点指向下一个节点。
	for list1!=nil && list2!= nil {
		// 如果左链表大，添加右链表节点
		if list1.Val >= list2.Val {
			p.Next = list2
			// 移动右链表
			list2 = list2.Next
		} else {
			p.Next = list1
			list1 = list1.Next
		}
		p = p.Next
	}
	//合并后 l1 和 l2 最多只有一个还未被合并完，我们直接将链表末尾指向未合并完的链表即可
	if list1 == nil {
		p.Next = list2
	} else {
		p.Next = list1
	}
	return head.Next
}
