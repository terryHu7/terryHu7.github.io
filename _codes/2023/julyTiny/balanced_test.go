package main

import "testing"

func TestBalancedBT(t *testing.T) {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var res bool

func isBalanced(root *TreeNode) bool {
	res = true
	if root == nil {
		return res
	}
	getDepthofNode(root)

	return res
}

/*
	 1     -> 1
 	/ \
 nil   nil -> 0
*/
func getDepthofNode(node *TreeNode) int {
	depth := 0
	if node == nil {
		return depth
	}

	ldepth := getDepthofNode(node.Left)
	rdepth := getDepthofNode(node.Right)

	if ldepth-rdepth > 1 || rdepth-ldepth > 1 {
		res = false
	}
	return max(ldepth, rdepth) + 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
