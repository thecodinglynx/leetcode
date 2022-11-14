package leetcode

import (
	"testing"

	ll "github.com/thecodinglynx/listutil"
)

func TestMiddleNode(t *testing.T) {
	var tests = []struct {
		list *ll.ListNode
		want *ll.ListNode
	}{
		{&ll.ListNode{Val: 1, Next: &ll.ListNode{Val: 2, Next: &ll.ListNode{Val: 3, Next: &ll.ListNode{Val: 4, Next: &ll.ListNode{Val: 5, Next: nil}}}}}, &ll.ListNode{Val: 3, Next: &ll.ListNode{Val: 4, Next: &ll.ListNode{Val: 5, Next: nil}}}},
	}
	for _, test := range tests {
		if got := middleNode(test.list); !ListsEqual(got, test.want) {
			t.Errorf("middleNode(%v) = %v - want %v", ll.ListToArr(test.list), ll.ListToArr(got), ll.ListToArr(test.want))
		}
	}
}

func ArrsEqual(l1, l2 []int) bool {
	if len(l1) != len(l2) {
		return false
	}
	for i := range l1 {
		if l1[i] != l2[i] {
			return false
		}
	}
	return true
}

func ListsEqual(l1, l2 *ll.ListNode) bool {
	return ArrsEqual(ll.ListToArr(l1), ll.ListToArr(l2))
}
