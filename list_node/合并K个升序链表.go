package main

import "math"

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	head := &ListNode{
		Val: math.MinInt,
	}
	dummy := head
	for _, list := range lists {
		dummy = mergeList(dummy, list)
	}
	return head.Next
}
func mergeList(list1, list2 *ListNode) *ListNode {
	dummy := &ListNode{}
	head := dummy
	for list1 != nil && list2 != nil {
		if list1.Val <= list2.Val {
			dummy.Next = list1
			list1 = list1.Next
		} else {
			dummy.Next = list2
			list2 = list2.Next
		}
		dummy = dummy.Next
	}
	if list1 == nil {
		dummy.Next = list2
	}
	if list2 == nil {
		dummy.Next = list1
	}
	return head.Next
}
