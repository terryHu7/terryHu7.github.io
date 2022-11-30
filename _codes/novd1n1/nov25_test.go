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

func uniquePathsII(obstacleGrid [][]int) int {
	//dp[i][j], (i, j) is the coordinate, dp[i][j] is the total number of paths
	//At the beginning, we tried to use obstacleGrid as the dp table, but it'll be a mess

	// initialize 2D slice
	m, n := len(obstacleGrid), len(obstacleGrid[0])
	dp := make([][]int, m)
	for index := range dp {
		dp[index] = make([]int, n)
	}
	okrow, okcolumn := true, true
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if obstacleGrid[0][j] != 1 {
				if okcolumn {
					dp[0][j] = 1
				}
			} else {
				okcolumn = false
			}
			if obstacleGrid[i][0] != 1 {
				if okrow {
					dp[i][0] = 1
				}
			} else {
				okrow = false
			}
		}
	}

	// if obstacleGrid[i][j]==1, correspondingly, dp[i][j] remains unchanged at 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if obstacleGrid[i][j] != 1 {
				dp[i][j] = dp[i][j-1] + dp[i-1][j]
			}
		}
	}
	return dp[m-1][n-1]
}

/*
func uniquePathsIIV2(obstacleGrid [][]int) int {
	//obstacleGrid[i][j], (i, j) is the coordinate, obstacleGrid[i][j] is the total number of paths
	for i := 0; i < len(obstacleGrid); i++ {
		for j := 0; j < len(obstacleGrid[0]); j++ {
			if i == 0 || j == 0 {
				// if i == 0 && obstacleGrid[i][j] == 1 {
				// 	obstacleGrid[0][j] = 0
				// }
				if obstacleGrid[i][j] == 1 { //obstacle
					obstacleGrid[i][j] = 0
				} else {
					obstacleGrid[i][j] = 1
				}
			} else {
				if obstacleGrid[i][j] == 1 {
					obstacleGrid[i][j] = 0
					continue
				}
				obstacleGrid[i][j] = obstacleGrid[i][j-1] + obstacleGrid[i-1][j]
			}

		}
	}
	return obstacleGrid[len(obstacleGrid)-1][len(obstacleGrid[0])-1]
}
*/
