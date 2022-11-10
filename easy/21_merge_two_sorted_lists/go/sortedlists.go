package leetcode

import (
	ll "github.com/thecodinglynx/listutil"
)

func mergeTwoLists(l1 *ll.ListNode, l2 *ll.ListNode) *ll.ListNode {
	pt := &ll.List{}
	for {
		if l1 != nil && l2 != nil {
			if l1.Val < l2.Val {
				pt.Insert(l1.Val)
				l1 = l1.Next
			} else {
				pt.Insert(l2.Val)
				l2 = l2.Next
			}
		} else if l1 != nil {
			pt.Insert(l1.Val)
			l1 = l1.Next
		} else if l2 != nil {
			pt.Insert(l2.Val)
			l2 = l2.Next
		} else {
			return pt.Head
		}
	}
}
