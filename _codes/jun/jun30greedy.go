package main

import "fmt"

func main() {
	// m := [...]int{
	// 	'a': 1,
	// 	'b': 2,
	// 	'c': 3,
	// }
	// m['a'] = 3
	// fmt.Printf("m:%T,%+v\n", m, m)
	// fmt.Println(m[0], len(m))
	// fmt.Println([...]int{1: 4, 5})
	fmt.Println(maxProfit([]int{7, 1, 8, 2, 9}))    //8
	fmt.Println(maxProfit([]int{7, 1, 5, 3, 6, 4})) //5
	fmt.Println(maxProfit([]int{7, 6, 4, 3, 1}))    //0
	fmt.Println(maxProfit([]int{1, 2, 3, 4, 5}))    //4
}

func maxProfit(prices []int) int {
	// if len(prices) <= 1 {
	// 	return 0
	// }
	s, max := 0, 0
	for i := 1; i < len(prices); i++ {
		if s >= 0 {
			s += (prices[i] - prices[i-1])
			if s < 0 {
				s = 0
			} else if s > max {
				max = s
			}
		}
	}

	return max
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) []int {
	var res []int
	if root == nil {
		return res
	}
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			node := queue[i]
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
			res = append(res, node.Val)
		}
		queue = queue[size:]
	}
	return res
}
