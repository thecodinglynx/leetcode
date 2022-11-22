package leetcode

import (
	ll "github.com/thecodinglynx/listutil"
)

// recursive solution
func reverseList(head *ll.ListNode) *ll.ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	return reverse(head)
}

func reverse(l *ll.ListNode) *ll.ListNode {
	if l.Next == nil {
		return l
	}
	cur := reverse(l.Next)
	insert(l.Val, cur)
	return cur
}

func insert(val int, l *ll.ListNode) {
	for l.Next != nil {
		l = l.Next
	}
	l.Next = &ll.ListNode{Val: val}
}
