package main

import (
	"math"
	"testing"
)

func Test18(t *testing.T) {
	// fmt.Println(maxPathSum(&TreeNode{1, &TreeNode{2, nil, nil}, &TreeNode{3, nil, nil}}))
	// fmt.Println(math.MinInt64 - 3)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var maxSum int

func maxPathSum(root *TreeNode) int {
	maxSum = math.MinInt32
	_ = traverse(root)
	return maxSum
}

func traverse(node *TreeNode) int {
	if node == nil {
		return math.MinInt32
	}

	leftMax := traverse(node.Left)
	rightMax := traverse(node.Right)
	maxSum = max(maxSum, leftMax, rightMax, node.Val,
		node.Val+leftMax, node.Val+rightMax,
		node.Val+leftMax+rightMax)
	return max(max(leftMax, rightMax)+node.Val, node.Val)
}

func max(nums ...int) int {
	res := nums[0]
	for _, num := range nums {
		if num > res {
			res = num
		}
	}
	return res
}
