package leetcode

import "testing"

func TestInOrder(t *testing.T) {
	var tests = []struct {
		root *TreeNode
		want []int
	}{
		{nil, []int{}},
		{&TreeNode{Val: 1}, []int{1}},
		{&TreeNode{Val: 3, Left: &TreeNode{Val: 9}}, []int{9, 3}},
		{&TreeNode{Val: 3, Left: &TreeNode{Val: 9}, Right: &TreeNode{Val: 20, Left: &TreeNode{Val: 15}, Right: &TreeNode{Val: 7}}}, []int{9, 3, 15, 20, 7}},
	}
	for _, test := range tests {
		if got := inorderTraversal(test.root); !ArrEqual(got, test.want) {
			t.Errorf("inorderTraversal(%v) = %v, want %v", test.root, got, test.want)
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
