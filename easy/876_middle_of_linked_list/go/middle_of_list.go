package leetcode

import (
	ll "github.com/thecodinglynx/listutil"
)

func middleNode(head *ll.ListNode) *ll.ListNode {
	if head.Next == nil {
		return head
	}
	if head.Next.Next == nil {
		return head.Next
	}

	p1 := head
	p2 := head.Next.Next

	for p2 != nil && p2.Next != nil {
		p2 = p2.Next.Next
		p1 = p1.Next
	}
	return p1.Next
}
