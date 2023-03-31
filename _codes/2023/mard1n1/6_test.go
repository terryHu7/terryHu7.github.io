package mard1n1

import "testing"

func TestTreeIteration(t *testing.T) {

}

func preorderTraversal_i(root *TreeNode) []int {
	res := []int{}
	if root == nil {
		return res
	}

	stack := []*TreeNode{root}
	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if node == nil {
			node = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			res = append(res, node.Val)
		} else {
			// preorder: middle->left->right
			// push: right->left->middle
			if node.Right != nil {
				stack = append(stack, node.Right)
			}
			if node.Left != nil {
				stack = append(stack, node.Left)
			}
			stack = append(stack, node)
			stack = append(stack, nil)
		}
	}
	return res
}

func inorderTraversal_i(root *TreeNode) []int {
	res := []int{}
	if root == nil {
		return res
	}

	stack := []*TreeNode{root}
	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if node == nil {
			node = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			res = append(res, node.Val)
		} else {
			// preorder: left->middle->right
			// push: right->middle->left
			if node.Right != nil {
				stack = append(stack, node.Right)
			}
			stack = append(stack, node)
			stack = append(stack, nil)
			if node.Left != nil {
				stack = append(stack, node.Left)
			}

		}
	}
	return res
}
