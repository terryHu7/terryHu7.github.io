package main

import (
	"fmt"
	"math"
	"testing"
)

func TestChange(t *testing.T) {
	fmt.Println(change(4, []int{1, 2, 3}))
	fmt.Println(combinationSum4([]int{1, 2, 3}, 4))
	fmt.Println(numSquares(12))
}

func change(amount int, coins []int) int {
	dp := make([]int, amount+1)
	dp[0] = 1 //like amount=1, coin=1

	for i := 0; i < len(coins); i++ {
		for j := coins[i]; j <= amount; j++ {
			dp[j] += dp[j-coins[i]]
		}
	}
	return dp[amount]
}

func combinationSum4(nums []int, target int) int {
	dp := make([]int, target+1)
	dp[0] = 1

	for i := 1; i <= target; i++ {
		for _, n := range nums {
			if n <= i {
				dp[i] += dp[i-n]
			}
		}
	}
	return dp[target]
}

func numSquares(n int) int {
	min := func(a, b int) int {
		if a <= b {
			return a
		}
		return b
	}

	dp := make([]int, n+1)
	for i := 1; i < len(dp); i++ {
		dp[i] = math.MaxInt32 - 1
	}

	for i := 1; i < len(dp); i++ {
		for j := 1; j*j <= i; j++ {
			dp[i] = min(dp[i], dp[i-j*j]+1)
		}
	}
	return dp[n]
}
