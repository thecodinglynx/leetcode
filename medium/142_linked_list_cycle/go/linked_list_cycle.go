package leetcode

import (
	ll "github.com/thecodinglynx/listutil"
)

func detectCycle(head *ll.ListNode) *ll.ListNode {
	m := make(map[*ll.ListNode]bool)
	for head != nil {
		if _, ok := m[head]; ok {
			return head
		}
		m[head] = true
		head = head.Next
	}
	return head
}
