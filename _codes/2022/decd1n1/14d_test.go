package main

import (
	"fmt"
	"testing"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func TestDPTree(t *testing.T) {
	fmt.Println(rob(&TreeNode{0, nil, nil}))
}
func rob(root *TreeNode) int {
	dp := traversal(root)
	return max(dp[0], dp[1])
}

func traversal(node *TreeNode) []int {
	if node == nil {
		return []int{0, 0}
	}
	dpL := traversal(node.Left)
	dpR := traversal(node.Right)
	return []int{node.Val + dpL[1] + dpR[1], max(dpL[0], dpL[1]) + max(dpR[0], dpR[1])}
}
