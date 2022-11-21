package leetcode

import (
	"testing"

	ll "github.com/thecodinglynx/listutil"
)

func TestRemoveFromList(t *testing.T) {
	var tests = []struct {
		input *ll.ListNode
		n     int
		want  *ll.ListNode
	}{

		{&ll.ListNode{Val: 1, Next: &ll.ListNode{Val: 2, Next: &ll.ListNode{Val: 3, Next: &ll.ListNode{Val: 4, Next: &ll.ListNode{Val: 5, Next: nil}}}}},
			2,
			&ll.ListNode{Val: 1, Next: &ll.ListNode{Val: 2, Next: &ll.ListNode{Val: 3, Next: &ll.ListNode{Val: 5, Next: nil}}}}},

		{&ll.ListNode{Val: 1, Next: nil}, 1, nil},

		{&ll.ListNode{Val: 1, Next: &ll.ListNode{Val: 2, Next: nil}},
			1,
			&ll.ListNode{Val: 1, Next: nil}},

		{&ll.ListNode{Val: 1, Next: &ll.ListNode{Val: 2, Next: nil}},
			2,
			&ll.ListNode{Val: 2, Next: nil}},
	}
	for _, test := range tests {
		if got := removeNthFromEnd(test.input, test.n); !ListsEqual(got, test.want) {
			t.Errorf("removeNthFromEnd(%v, %d) = %v, want %v", ll.ListToArr(test.input), test.n, ll.ListToArr(got), ll.ListToArr(test.want))
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
