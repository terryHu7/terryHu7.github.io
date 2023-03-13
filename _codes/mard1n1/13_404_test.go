package mard1n1

import "testing"

func Test404(t *testing.T) {

}

var res int

func sumOfLeftLeaves(root *TreeNode) int {
	res := 0
	traverse404(root)
	return res
}

func traverse404(node *TreeNode) {
	if node == nil {
		return
	}
	if node.Left != nil && node.Left.Left == nil && node.Left.Right == nil {
		res += node.Left.Val
	}
	traverse404(node.Left)
	traverse404(node.Left)
}

var maxDeep int // 全局变量 深度
var value int   //全局变量 最终值
func findBottomLeftValue(root *TreeNode) int {
	if root.Left == nil && root.Right == nil { //需要提前判断一下（不要这个if的话提交结果会出错，但执行代码不会。防止这种情况出现，故先判断是否只有一个节点）
		return root.Val
	}
	findLeftValue(root, maxDeep)
	return value
}
func findLeftValue(root *TreeNode, deep int) {
	//最左边的值在左边
	if root.Left == nil && root.Right == nil {
		if deep > maxDeep {
			value = root.Val
			maxDeep = deep
		}
	}
	//递归
	if root.Left != nil {
		deep++
		findLeftValue(root.Left, deep)
		deep-- //回溯
	}
	if root.Right != nil {
		deep++
		findLeftValue(root.Right, deep)
		deep-- //回溯
	}
}
