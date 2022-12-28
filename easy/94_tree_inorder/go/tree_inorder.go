package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	res := make([]int, 0)
	return visit(root, res)

}

func visit(root *TreeNode, arr []int) []int {
	if root == nil {
		return arr
	}
	arr = visit(root.Left, arr)
	arr = append(arr, root.Val)
	arr = visit(root.Right, arr)
	return arr
}
