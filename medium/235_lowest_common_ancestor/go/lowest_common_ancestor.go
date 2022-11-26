package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// finds the path for both p and q and stores them in a slice.
// Then iterates through the slice backwards to identify the common ancestor
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	pPtr := root
	qPtr := root
	pPath := []*TreeNode{}
	qPath := []*TreeNode{}
	for pPtr != nil {
		pPath = append(pPath, pPtr)
		if pPtr.Val == p.Val {
			break
		}
		if pPtr.Val > p.Val {
			pPtr = pPtr.Left
		} else {
			pPtr = pPtr.Right
		}

	}
	for qPtr != nil {
		qPath = append(qPath, qPtr)
		if qPtr.Val == q.Val {
			break
		}
		if qPtr.Val > q.Val {
			qPtr = qPtr.Left
		} else {
			qPtr = qPtr.Right
		}
	}

	for i := len(pPath) - 1; i >= 0; i-- {
		for j := len(qPath) - 1; j >= 0; j-- {
			if pPath[i].Val == qPath[j].Val {
				return pPath[i]
			}
		}
	}
	return nil
}
