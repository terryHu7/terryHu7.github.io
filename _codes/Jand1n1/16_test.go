package main

import (
	"fmt"
	"testing"
)

func TestSpiral(t *testing.T) {
	fmt.Println(generateMatrix(4))
}

func generateMatrix(n int) [][]int {
	left, right := 0, n-1
	top, bottom := 0, n-1

	num, target := 1, n*n

	matrix := make([][]int, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]int, n)
	}

	for num <= target {
		if num == target && n%2 == 1 {
			matrix[int(n/2)][int(n/2)] = target
			break
		}
		for i := left; i < right; i++ {
			matrix[top][i] = num
			num++
		}
		for i := top; i < bottom; i++ {
			matrix[i][right] = num
			num++
		}
		for i := right; i > left; i-- {
			matrix[bottom][i] = num
			num++
		}
		for i := bottom; i > top; i-- {
			matrix[i][left] = num
			num++
		}
		left++
		right--
		top++
		bottom--
	}
	return matrix
}

func spiralOrder(matrix [][]int) []int {
	m, n := len(matrix), len(matrix[0])
	left, right := 0, n-1
	top, bottom := 0, m-1

	num, target := 1, m*n

	res := []int{}
	for num <= target {
		if num == target && m == n && n%2 == 1 {
			res = append(res, matrix[top][left])
			break
		}
		for i := left; i < right && num <= target; i++ {
			res = append(res, matrix[top][i])
			num++
		}
		for i := top; i < bottom && num <= target; i++ {
			res = append(res, matrix[i][right])
			num++
		}
		for i := right; i > left && num <= target; i-- {
			res = append(res, matrix[bottom][i])
			num++
		}
		for i := bottom; i > top && num <= target; i-- {
			res = append(res, matrix[i][left])
			num++
		}
		left++
		right--
		top++
		bottom--
	}
	return res
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeElements(head *ListNode, val int) *ListNode {
	dummy := &ListNode{}
	dummy.Next = head

	cur := dummy
	for cur != nil {
		if cur.Next != nil && cur.Next.Val == val {
			cur.Next = cur.Next.Next
		}
		cur = cur.Next
	}
	return dummy.Next
}
