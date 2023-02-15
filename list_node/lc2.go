package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// 看答案了
func addTwoNumbers(l1 *ListNode, l2 *ListNode) (ans *ListNode) {
	var tail *ListNode
	carry := 0
	// 走完最长的那个，短的缺失可以理解为0
	for l1 != nil || l2 != nil {
		n1, n2 := 0, 0
		if l1 != nil {
			n1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			n2 = l2.Val
			l2 = l2.Next
		}
		sum := n1 + n2 + carry
		// 设置的sum不能超过10
		sum, carry = sum%10, sum/10
		if ans == nil {
			// 构造头节点
			ans = &ListNode{
				Val: sum,
			}
			tail = ans
		} else {
			// tail后移一位
			tail.Next = &ListNode{Val: sum}
			tail = tail.Next
		}
	}
	if carry > 0 {
		// 最后一位需要增加一个节点
		tail.Next = &ListNode{Val: carry}
	}
	return
}
