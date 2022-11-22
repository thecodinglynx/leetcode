package leetcode

import (
	"testing"

	ll "github.com/thecodinglynx/listutil"
)

func TestRemoveFromList(t *testing.T) {
	var tests = []struct {
		input *ll.ListNode
		want  *ll.ListNode
	}{
		{&ll.ListNode{Val: 1, Next: &ll.ListNode{Val: 2, Next: &ll.ListNode{Val: 3, Next: &ll.ListNode{Val: 4, Next: &ll.ListNode{Val: 5, Next: nil}}}}},
			&ll.ListNode{Val: 5, Next: &ll.ListNode{Val: 4, Next: &ll.ListNode{Val: 3, Next: &ll.ListNode{Val: 2, Next: &ll.ListNode{Val: 1, Next: nil}}}}}},
	}
	for _, test := range tests {
		if got := reverseList(test.input); !ListsEqual(got, test.want) {
			t.Errorf("reverseList(%v) = %v, want %v", ll.ListToArr(test.input), ll.ListToArr(got), ll.ListToArr(test.want))
		}
	}
}

func ArrEqual(a1, a2 []int) bool {
	if len(a1) != len(a2) {
		return false
	}
	for i := range a1 {
		if a1[i] != a2[i] {
			return false
		}
	}
	return true
}

func ListsEqual(l1, l2 *ll.ListNode) bool {
	return ArrEqual(ll.ListToArr(l1), ll.ListToArr(l2))
}
