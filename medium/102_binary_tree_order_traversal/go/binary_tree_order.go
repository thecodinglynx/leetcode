package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	res := make([][]int, 0)
	return visit(root, 0, res)
}

func visit(root *TreeNode, lvl int, arr [][]int) [][]int {
	if root == nil {
		return arr
	}
	if len(arr) <= lvl {
		arr = append(arr, []int{})
	}
	arr[lvl] = append(arr[lvl], root.Val)
	lvl++
	arr = visit(root.Left, lvl, arr)
	arr = visit(root.Right, lvl, arr)
	return arr
}
