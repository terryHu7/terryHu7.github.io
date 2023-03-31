package mard1n1

import "testing"

func TestBSTInsert(t *testing.T) {

}

func insertIntoBST(root *TreeNode, val int) *TreeNode {
	insertIntoBSTHelper(root, val)
	return root
}

func insertIntoBSTHelper(node *TreeNode, val int) {
	if node == nil {
		return
	}
	if node.Left != nil {
		_ = insertIntoBST(node.Left, val)
	}
	if node != nil {
		_ = insertIntoBST(node.Right, val)
	}

	if node.Left == nil && val < node.Val {
		node.Left = &TreeNode{
			Val: val,
		}
	}
	if node.Right == nil && val > node.Val {
		node.Right = &TreeNode{
			Val: val,
		}
	}
	return

}
