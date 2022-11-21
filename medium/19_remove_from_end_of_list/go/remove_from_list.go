package leetcode

import (
	ll "github.com/thecodinglynx/listutil"
)

func removeNthFromEnd(head *ll.ListNode, n int) *ll.ListNode {
	len := 0
	root := head
	for head != nil {
		len++
		head = head.Next
	}
	res := &ll.List{}
	nSkip := len - n
	c := 0
	for root != nil {
		if c == nSkip {
			root = root.Next
		}
		if root != nil {
			res.Insert(root.Val)
			root = root.Next
		}
		c++
	}
	return res.Head
}
