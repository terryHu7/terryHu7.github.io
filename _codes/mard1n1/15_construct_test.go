package mard1n1

import "testing"

func TestConstruct(t *testing.T) {

}

func constructMaximumBinaryTree(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	index := maxValAndIndex(nums)
	node := &TreeNode{
		Val:   nums[index],
		Left:  &TreeNode{},
		Right: &TreeNode{},
	}

	node.Left = constructMaximumBinaryTree(nums[:index])
	node.Right = constructMaximumBinaryTree(nums[index+1:])
	return node
}

func maxValAndIndex(nums []int) int {
	if len(nums) == 0 {
		return -1
	}
	max := nums[0]
	res := 0

	for i := 0; i < len(nums); i++ {
		if nums[i] > max {
			max = nums[i]
			res = i
		}
	}
	return res
}

func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(postorder) == 0 { // preorder[0] & postorder[len(postorder)-1] is the root
		return nil
	}
	index := findIndex(inorder, postorder)
	left := buildTree(inorder[:index], postorder[:index]) //same len
	right := buildTree(inorder[index+1:], postorder[index:len(postorder)-1])
	return &TreeNode{
		Val:   postorder[len(postorder)-1],
		Left:  left,
		Right: right,
	}

}

//中序遍历 inorder = [9,3,15,20,7]
//后序遍历 postorder = [9,15,7,20,3]
func findIndex(inorder []int, postorder []int) int {
	rootVal := postorder[len(postorder)-1]
	for i, v := range inorder {
		if v == rootVal {
			return i
		}
	}
	return -1
}
