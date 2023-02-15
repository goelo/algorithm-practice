package main

// lc160,忘记了看官网题解证明部分
// 双指针问题
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	// 定义两个指针
	pa := headA
	pb := headB
	// 如果这个指针一直不相等
	for pa != pb {
		// 如果pa为空
		if pa == nil {
			// 则将pa指向b的头节点,
			// 链表 headA\textit{headA}headA 和 headB\textit{headB}headB 的长度分别是 mmm 和 nnn。假设链表 headA\textit{headA}headA 的不相交部分有 aaa 个节点，链表 headB\textit{headB}headB 的不相交部分有 bbb 个节点，两个链表相交的部分有 ccc 个节点，则有 a+c=ma+c=ma+c=m，b+c=nb+c=nb+c=n。
			// 这里之所以这样做，是a和b走到头后互换的话，最终指针如果存在相交节点，
			// 那第一次走到相交节点的步数，走的步数是一样
			pa = headB
		} else {
			// 否则pa后移
			pa = pa.Next
		}
		if pb == nil {
			pb = headA
		} else {
			pb = pb.Next
		}
	}
	return pb
}
