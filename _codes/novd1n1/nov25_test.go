package main

import (
	"fmt"
	"testing"
)

func TestUniquePaths(t *testing.T) {
	fmt.Println(uniquePaths(3, 4))
}

func uniquePaths(m, n int) int {
	//dp[i][j], (i, j) is the coordinate, dp[i][j] is the total number of paths

	// initialize 2D slice
	dp := make([][]int, m)
	for index := range dp {
		dp[index] = make([]int, n)
	}
	dp[0][0] = 1

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 || j == 0 {
				dp[i][j] = 1
			} else {
				dp[i][j] = dp[i][j-1] + dp[i-1][j]
			}
			if i == m-1 && j == n-1 {
				return dp[i][j]
			}
		}
	}
	return -1
}
