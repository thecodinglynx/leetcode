package leetcode

import "testing"

func TestLevelOrder(t *testing.T) {
	var tests = []struct {
		root *TreeNode
		want [][]int
	}{
		{nil, [][]int{}},
		{&TreeNode{Val: 1}, [][]int{{1}}},
		{&TreeNode{Val: 3, Left: &TreeNode{Val: 9}}, [][]int{{3}, {9}}},
		{&TreeNode{Val: 3, Left: &TreeNode{Val: 9}, Right: &TreeNode{Val: 20, Left: &TreeNode{Val: 15}, Right: &TreeNode{Val: 7}}}, [][]int{{3}, {9, 20}, {15, 7}}},
	}
	for _, test := range tests {
		if got := levelOrder(test.root); !ArrEqual(got, test.want) {
			t.Errorf("levelOrder(%v) = %v, want %v", test.root, got, test.want)
		}
	}
}

func ArrEqual(a1, a2 [][]int) bool {
	if len(a1) != len(a2) {
		return false
	}
	for i := range a1 {
		for j := range a1[i] {
			if a1[i][j] != a2[i][j] {
				return false
			}
		}
	}
	return true
}
