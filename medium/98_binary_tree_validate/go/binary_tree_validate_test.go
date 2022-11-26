package leetcode

import "testing"

func TestValidBST(t *testing.T) {
	var tests = []struct {
		root *TreeNode
		want bool
	}{
		{nil, true},
		{&TreeNode{Val: 1}, true},
		{&TreeNode{Val: 0}, true},
		{&TreeNode{Val: 3, Left: &TreeNode{Val: 9}}, false},
		{&TreeNode{Val: 2, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 3}}, true},
		{&TreeNode{Val: 3, Left: &TreeNode{Val: 9}, Right: &TreeNode{Val: 20, Left: &TreeNode{Val: 15}, Right: &TreeNode{Val: 7}}}, false},
		{&TreeNode{Val: 5, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 4, Left: &TreeNode{Val: 3}, Right: &TreeNode{Val: 6}}}, false},
		{&TreeNode{Val: 5, Left: &TreeNode{Val: 4}, Right: &TreeNode{Val: 6, Left: &TreeNode{Val: 3}, Right: &TreeNode{Val: 7}}}, false},
		{&TreeNode{Val: 4, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 3}}, Right: &TreeNode{Val: 6, Left: &TreeNode{Val: 5}, Right: &TreeNode{Val: 7}}}, true},
	}
	for _, test := range tests {
		if got := isValidBST(test.root); got != test.want {
			t.Errorf("isValidBST(%v) = %t", test.root, got)
		}
	}
}
