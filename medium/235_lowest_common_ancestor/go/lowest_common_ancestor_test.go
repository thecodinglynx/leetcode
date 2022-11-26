package leetcode

import "testing"

func TestLowestCommonAncestor(t *testing.T) {
	l1 := &TreeNode{Val: 6}
	l2_l := &TreeNode{Val: 2}
	l2_r := &TreeNode{Val: 8}
	l3_l_1 := &TreeNode{Val: 0}
	l3_r_1 := &TreeNode{Val: 4}
	l3_l_2 := &TreeNode{Val: 7}
	l3_r_2 := &TreeNode{Val: 9}
	l4_l := &TreeNode{Val: 3}
	l4_r := &TreeNode{Val: 5}

	root := l1
	root.Left = l2_l
	root.Right = l2_r
	l2_l.Left = l3_l_1
	l2_l.Right = l3_r_1
	l2_r.Left = l3_l_2
	l2_r.Right = l3_r_2
	l3_r_1.Left = l4_l
	l3_r_1.Right = l4_r

	if got := lowestCommonAncestor(root, l4_l, l3_r_1); got.Val != l3_r_1.Val {
		t.Errorf("lowestCommonAncestor(%v, %v, %v) = %v", root, l4_l, l3_r_1, got)
	}
}
