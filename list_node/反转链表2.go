package main

func reverseLink(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	// 注意这里声明空结构体需要注意会有pre的默认值0，pre的初始化必须为nil
	var pre *ListNode

	cur := head
	for cur != nil {
		// 临时记录下一个节点
		tmp := cur.Next
		// 将当前cur的节点的下一个节点反转
		cur.Next = pre
		// 移动pre和cur节点到下一个, 需要注意顺序
		pre = cur
		cur = tmp

	}
	return cur
}
