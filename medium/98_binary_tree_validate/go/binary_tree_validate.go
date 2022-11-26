package leetcode

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Res struct {
	valid bool
	val   int
}

// traverses the BST and ensures that it's properly sorted
func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}
	val := math.MinInt
	res := visit(root, val)
	return res.valid
}

func visit(node *TreeNode, val int) Res {
	if node == nil {
		return Res{valid: true, val: val}
	}
	res := visit(node.Left, val)
	if !res.valid {
		return Res{valid: false}
	}
	if node.Val <= res.val {
		return Res{valid: false}
	}
	val = node.Val
	return visit(node.Right, val)
}
