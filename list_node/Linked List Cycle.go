package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	// 这里这样做，是假设有一个虚拟头节点，然后slow走一步，fast走两部
	slow := head
	fast := head.Next
	// 标准解法
	//for fast != slow {
	//	// 到达尾部节点
	//	if fast == nil || fast.Next == nil {
	//		return false
	//	}
	//	slow = slow.Next
	//	fast = fast.Next.Next
	//}
	for slow != nil && fast != nil {
		// 如果fast到达尾部节点, 肯定无环
		if fast.Next == nil {
			return false
		}
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			return true
		}
	}
	return false
}
