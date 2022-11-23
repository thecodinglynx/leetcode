package leetcode

import (
	"testing"

	ll "github.com/thecodinglynx/listutil"
)

func TestLinkedListCycle(t *testing.T) {

	l1 := &ll.ListNode{Val: 3, Next: nil}
	l2 := &ll.ListNode{Val: 2, Next: nil}
	l3 := &ll.ListNode{Val: 0, Next: nil}
	l4 := &ll.ListNode{Val: -4, Next: nil}

	l1.Next = l2
	l2.Next = l3
	l3.Next = l4
	cycleStart := l3
	l4.Next = cycleStart

	if got := detectCycle(l1); got.Val != cycleStart.Val {
		t.Errorf("reverseList(%v...) = %v..., want %v...", l1.Val, got.Val, cycleStart.Val)
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
