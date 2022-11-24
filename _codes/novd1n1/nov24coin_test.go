package main

import (
	"fmt"
	"testing"
)

func TestCoin(t *testing.T) {
	fmt.Println(coinChange([]int{1, 2, 5}, 11))
}

func coinChange(coins []int, amount int) int {
	//dp[i], i stands for amount, dp[i] stands for the num of coins
	dp := make([]int, amount+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = amount + 1 // act like math.MaxInt
	}
	dp[0] = 0
	for i := 1; i < len(dp); i++ {
		for c := range coins {
			if i < c {
				continue
			}
			dp[i] = minimum(dp[i], dp[i-c]+1)
		}
	}
	if dp[amount] == amount+1 {
		return -1
	} else {
		return dp[amount]
	}
}

func minimum(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
