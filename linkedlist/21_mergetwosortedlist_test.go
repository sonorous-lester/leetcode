package linkedlist

import (
	"testing"
)

/*
[(3) Merge Two Sorted Lists - LeetCode](https://leetcode.com/problems/merge-two-sorted-lists/)
*/

func TestMergeTwoSortedList(t *testing.T) {
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	var head ListNode
	cur := &head

	for list1 != nil && list2 != nil {
		if list1.Val <= list2.Val {
			cur.Next = list1
			list1 = list1.Next
		} else {
			cur.Next = list2
			list2 = list2.Next
		}
		cur = cur.Next
	}

	if list1 != nil {
		cur.Next = list1
	} else {
		cur.Next = list2
	}
	return head.Next
}
