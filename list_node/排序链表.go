package main

// 自顶而上
func SortList(head *ListNode) *ListNode {
	// ！！！这里终止条件判断两种
	if head == nil || head.Next == nil {
		return head
	}
	// 找到中间位置, 用快慢指针
	// slow是指向的是中间位置
	slow, fast := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	// 从中间结点截断
	right := slow.Next
	slow.Next = nil
	// 左半边
	l := SortList(head)
	// 右半边
	r := SortList(right)
	// 合并左右链表
	return merge(l, r)
}

func merge1(head1, head2 *ListNode) *ListNode {
	var dummy *ListNode
	cur := dummy
	p1, p2 := head1, head2
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

// 自底而上 O(1) 空间复杂度
func sortList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	// 链表长度
	length := 0
	// 统计链表长度
	for node := head; node != nil; node = node.Next {
		length++
	}
	// 定义虚拟指针，指向head
	dummyHead := &ListNode{Next: head}
	// 这里注意从1开始，subLength表示每次需要排序的链表长度
	// 每次合并之后的链表长度都是之前的两倍，因此subLength*=2
	for subLength := 1; subLength < length; subLength <<= 1 {
		prev, cur := dummyHead, dummyHead.Next
		// 这个循环相当于把链表从头开始分割成单个节点的链表，然后再合并成两两，四四的链表
		for cur != nil {
			head1 := cur
			// 每次要排序的链表不停后移,第一次不会进入这个循环。因为只有一个节点，跳过
			for i := 1; i < subLength && cur.Next != nil; i++ {
				cur = cur.Next
			}
			// 第一次的时候head2是第二个节点
			head2 := cur.Next
			// 截断链表,head1变成单一节点链表
			cur.Next = nil
			// cur移动到第二个节点,第一次循环的时候
			cur = head2
			// 第一次也不会进入这个循环
			for i := 1; i < subLength && cur != nil && cur.Next != nil; i++ {
				cur = cur.Next
			}

			var next *ListNode
			if cur != nil {
				next = cur.Next
				// 截断head2后续的链表，head2也变成了单一节点链表
				cur.Next = nil
			}
			// 合并这两个单一节点链表链表
			prev.Next = merge(head1, head2)
			// 合并完之后，prev移动到尾部
			for prev.Next != nil {
				prev = prev.Next
			}
			// 移动cur指针，后移到连续的链表，截断的链表之后还是连续，继续从头开始
			cur = next
		}
	}
	return dummyHead.Next
}

// 合并两个有序链表
func merge(head1, head2 *ListNode) *ListNode {
	// 定义空节点
	dummyHead := &ListNode{}
	temp, p1, p2 := dummyHead, head1, head2
	// 只要链表1和链表2都没走到尾部
	for p1 != nil && p2 != nil {
		// 如果左链表小于右链表，左链表链接到dummy上, temp1后移
		if p1.Val <= p2.Val {
			temp.Next = p1
			p1 = p1.Next
		} else {
			temp.Next = p2
			p2 = p2.Next
		}
		// 链接完之后temp后移
		temp = temp.Next
	}
	if p1 != nil {
		temp.Next = p1
	} else if p2 != nil {
		temp.Next = p2
	}
	return dummyHead.Next
}
