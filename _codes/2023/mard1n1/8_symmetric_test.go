package mard1n1

import (
	"testing"
)

func TestSymmetric(t *testing.T) {

}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return traverse(root.Left, root.Right)
}

func traverse(n1, n2 *TreeNode) bool {
	res := false
	switch {
	case n1 == nil && n2 == nil:
		res = true
	case n1 != nil && n2 != nil && n1.Val == n2.Val:
		res = true
		outer := traverse(n1.Left, n2.Right)
		inner := traverse(n1.Right, n2.Left)
		res = outer && inner
	}
	return res
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	res := false

	switch {
	case p == nil && q == nil:
		res = true
	case p != nil && q != nil && p.Val == q.Val:
		res = true
		lres := isSameTree(p.Left, q.Left)
		rres := isSameTree(p.Right, q.Right)
		res = lres && rres
	}
	return res
}
