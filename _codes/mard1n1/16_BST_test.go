package mard1n1

import "testing"

func TestBST(t *testing.T) {

}

func searchBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return nil
	}

	if root.Val > val {
		return searchBST(root.Left, val)
	} else if root.Val < val {
		return searchBST(root.Right, val)
	} else {
		return root
	}
}

var iarray []int

func isValidBST(root *TreeNode) bool {
	iarray = []int{}
	inorderTraverse(root)
	for i := 1; i < len(iarray); i++ {
		if iarray[i] <= iarray[i-1] {
			return false
		}
	}
	return true
}

func inorderTraverse(node *TreeNode) {
	if node == nil {
		return
	}
	inorderTraverse(node.Left)
	iarray = append(iarray, node.Val)
	inorderTraverse(node.Right)
}
