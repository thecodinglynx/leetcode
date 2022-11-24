package leetcode

import "testing"

func TestPreorder(t *testing.T) {
	var tests = []struct {
		root *Node
		want []int
	}{
		{nil, []int{}},
		{&Node{Val: 1}, []int{1}},
		{&Node{Val: 1, Children: []*Node{&Node{Val: 2}}}, []int{1, 2}},
		{&Node{Val: 1, Children: []*Node{&Node{Val: 3, Children: []*Node{&Node{Val: 5}, &Node{Val: 6}}}, &Node{Val: 2}, &Node{Val: 4}}}, []int{1, 3, 5, 6, 2, 4}},
	}
	for _, test := range tests {
		if got := preorder(test.root); !ArrEqual(got, test.want) {
			t.Errorf("preorder(%v) = %v, want %v", test.root, got, test.want)
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
