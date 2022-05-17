package linkedlist

import (
	"testing"
)

/*
[(3) Linked List Cycle - LeetCode](https://leetcode.com/problems/linked-list-cycle/)
*/

func TestLinkedListCycle(t *testing.T) {
}

func hasCycle(head *ListNode) bool {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow, fast = slow.Next, fast.Next.Next
		if slow == fast {
			return true
		}
	}
	return false
}
