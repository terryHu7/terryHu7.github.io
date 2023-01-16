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
