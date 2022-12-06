package main

import (
	"fmt"
	"testing"
)

func Test01Package(t *testing.T) {
	fmt.Println(canPartition([]int{1, 5, 11, 5}))
}

func canPartition(nums []int) bool {
	//sum/2
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}
	if sum%2 != 0 {
		return false
	}
	target := sum / 2
	//init
	dp := make([][]int, len(nums))
	for i := range dp {
		dp[i] = make([]int, target+1)
	}
	for j := nums[0]; j <= target; j++ {
		dp[0][j] = nums[0]
	}
	//dp
	for i := 1; i < len(nums); i++ {
		for j := 0; j <= target; j++ {
			if j >= nums[i] {
				dp[i][j] = max(
					dp[i-1][j],
					dp[i-1][j-nums[i]]+nums[i])
			} else {
				dp[i][j] = dp[i-1][j]
			}
		}
	}
	return dp[len(nums)-1][target] == target
}
func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
