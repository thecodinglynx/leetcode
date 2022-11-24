package leetcode

type Node struct {
	Val      int
	Children []*Node
}

func preorder(root *Node) []int {
	res := make([]int, 0)
	return traverse(root, res)
}

func traverse(n *Node, arr []int) []int {
	if n == nil {
		return arr
	}
	arr = append(arr, n.Val)
	for c := range n.Children {
		arr = traverse(n.Children[c], arr)
	}
	return arr
}
