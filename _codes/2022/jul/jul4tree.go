package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	treetest := TreeNode{}
	fmt.Println(treetest.Val)
}

func tree_simplification(root *TreeNode) {

	all := []*TreeNode{root}
	layer := []*TreeNode{root}
	for _, n := range layer {
		layerlen := len(layer)
		if n.Left != nil {
			all = append(all, n.Left)
			layer = append(layer, n.Left)
		}
		if n.Right != nil {
			all = append(all, n.Right)
			layer = append(layer, n.Right)
		}
		layer = layer[layerlen:] //new layer
	}

}
