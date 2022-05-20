package main

import "fmt"

func main() {
	fmt.Println(robotdp(3, 4))
}

func robotdp(row, column int) (ways int) {
	//2d slice
	dp := [][]int{}
	for i := 0; i < row; i++ {
		dp = append(dp, make([]int, column))
	}

	for i := 0; i < row; i++ {
		for j := 0; j < column; j++ {
			if i == 0 || j == 0 {
				dp[i][j] = 1
			} else {
				dp[i][j] = dp[i-1][j] + dp[i][j-1]
			}
		}
	}

	return dp[row-1][column-1]
}
