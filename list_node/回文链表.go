package main

// 整个流程可以分为以下五个步骤：
// 1. 找到前半部分链表的尾节点。
// 2. 反转后半部分链表。
// 3. 判断是否回文。
// 4. 恢复链表。
// 5. 返回结果。
func isPalindrome(head *ListNode) bool {
	if head == nil {
		return true
	}
	// 1. 找到前半部分链表的尾结点
	first := FirstHalf(head)
	// 2. 反转后半部分的链表
	second := reverse(first.Next)
	// 3. 判断是否是回文
	p1 := head
	p2 := second
	res := true
	// 这里画图看一下，p2要短一些
	for p2 != nil && res {
		if p1.Val != p2.Val {
			return false
		}
		p1 = p1.Next
		p2 = p2.Next
	}
	// 恢复链表
	first.Next = reverse(second)
	return true
}

// 这里的快慢指针都是从头结点开始，如果从虚拟结点开始，那么奇数链表的时候，找不到前半部分。要注意
// 另外要注意终止条件，判断的是快指针的next以及next.next不为空就停止。这样就覆盖了奇数链表和偶数链表
func FirstHalf(head *ListNode) *ListNode {
	slow, fast := head, head
	// 这里的判断条件可以覆盖奇数和偶数链表
	for fast.Next != nil && fast.Next.Next != nil {
		// 慢指针一次走一步，快指针一次走两步
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

func reverse(head *ListNode) *ListNode {
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
