package main

import (
	"fmt"
	"testing"
)

func TestWithFee(t *testing.T) {
	fmt.Println(maxProfitwithFee([]int{1, 3, 2, 8, 4, 9}, 2))
}

func maxProfitwithFee(prices []int, fee int) int {
	dp := make([]int, 2)
	dp[0], dp[1] = -prices[0], 0

	for i := 1; i < len(prices); i++ {
		last0, last1 := dp[0], dp[1]
		dp[0] = max(last0, last1-prices[i])
		dp[1] = max(last1, last0+prices[i]-fee)
	}

	return dp[1]
}
