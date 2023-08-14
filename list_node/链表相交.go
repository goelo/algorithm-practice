package main

func getIntersectionNode1(headA, headB *ListNode) *ListNode {
	// 题目给出了A和B肯定不是空链表，所以不判断空了
	preA := &ListNode{
		Next: headA,
	}
	preB := &ListNode{
		Next: headB,
	}
	var (
		m int
		n int
	)
	for preA.Next != nil {
		m++
		preA = preA.Next
	}
	for preB.Next != nil {
		n++
		preB = preB.Next
	}

	var (
		fast *ListNode
		slow *ListNode
		step int
	)
	// 判断A B哪个长
	if m > n {
		step = m - n
		fast, slow = headA, headB

	} else {
		step = n - m
		fast, slow = headB, headA
	}
	for i := step; i > 0; i-- {
		fast = fast.Next
	}

	for fast != slow {
		fast = fast.Next
		slow = slow.Next
	}
	return fast
}

// 双指针方法，会让代码更简洁
func getIntersectionNode2(headA, headB *ListNode) *ListNode {
	var (
		p1 *ListNode
		p2 *ListNode
	)
	// 如果不想等，就一直移动指针
	p1 = headA
	p2 = headB
	for p1 != p2 {
		// 如果A不是空，就移动
		if p1 != nil {
			p1 = p1.Next
		} else {
			p1 = headB
		}
		if p2 != nil {
			p2 = p2.Next
		} else {
			p2 = headA
		}
	}
	return p1
}
