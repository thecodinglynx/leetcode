package leetcode

import (
	"testing"

	ll "github.com/thecodinglynx/listutil"
)

func TestMergeTwoLists(t *testing.T) {
	var tests = []struct {
		l1   []int
		l2   []int
		want *ll.ListNode
	}{
		{[]int{}, []int{}, nil},
		{[]int{5, 6, 7}, []int{},
			&ll.ListNode{Val: 5, Next: &ll.ListNode{Val: 6, Next: &ll.ListNode{Val: 7, Next: nil}}}},
		{[]int{5, 6, 7}, []int{5, 6},
			&ll.ListNode{Val: 5, Next: &ll.ListNode{Val: 5, Next: &ll.ListNode{Val: 6, Next: &ll.ListNode{Val: 6, Next: &ll.ListNode{Val: 7, Next: nil}}}}}},
	}
	for _, test := range tests {
		if got := mergeTwoLists(ll.ArrToList(test.l1), ll.ArrToList(test.l2)); !ListsEqual(got, test.want) {
			t.Errorf("mergeTwoLists(%v, %v) = %v - want %v", test.l1, test.l2, ll.ListToArr(got), ll.ListToArr(test.want))
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
