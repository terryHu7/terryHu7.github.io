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

func TestBTree(t *testing.T) {
	fmt.Println(minCameraCover(&TreeNode{
		Val: 0,
		Left: &TreeNode{
			Val:   0,
			Left:  nil,
			Right: nil,
		},
		Right: &TreeNode{
			Val:   0,
			Left:  nil,
			Right: nil,
		},
	}))
}

type TreeNodeWithTag struct {
	*TreeNode
	Tag bool
}

var res int

func minCameraCover(root *TreeNode) int {
	res = 0
	traverse(&TreeNodeWithTag{root, true}) //special case, for root, it's always not gonna ++
	if res == 0 && root != nil {
		return 1
	}
	return res
}

func traverse(node *TreeNodeWithTag) {
	if node.TreeNode == nil {
		return
	}

	tag := false
	if !node.Tag {
		res++
		tag = true
	}
	traverse(&TreeNodeWithTag{node.Left, tag})
	traverse(&TreeNodeWithTag{node.Right, tag})
}
