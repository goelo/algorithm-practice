package main

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	// 这种类似差值的，两个指针。快指针直接走n步，慢指针再开始走。当快指针走到结尾，慢指针指向的就是要删除的前一节点。
	// 注意这里的哨兵节点，很多题都是用哨兵节点。并不是从头指针开始
	p := &ListNode{}
	fast, slow := p, p
	p.Next = head
	for i := n; i > 0; i-- {
		fast = fast.Next
	}
	// 当快指针走到尾部时
	for fast.Next != nil {
		fast = fast.Next
		slow = slow.Next
	}
	// 删除指定节点
	slow.Next = slow.Next.Next
	return p.Next
}
