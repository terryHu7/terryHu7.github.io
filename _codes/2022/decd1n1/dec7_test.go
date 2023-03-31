package main

import (
	"fmt"
	"math"
	"testing"
)

func TestDP(t *testing.T) {
	fmt.Println(findTargetSumWays([]int{0, 0}, 0))
}

func findTargetSumWays(nums []int, S int) int {
	//sum/2
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}
	S = int(math.Abs(float64(S)))

	target := (sum + S) / 2
	if (sum+S)%2 != 0 || math.Abs(float64(target)) > math.Abs(float64(sum)) {
		return 0
	}

	//init if nums[i]=j 1
	dp := make([][]int, len(nums)+1)
	for i := range dp {
		dp[i] = make([]int, target+1)
	}
	//j can be 0
	dp[0][0] = 1
	for i := 1; i < len(nums)+1; i++ {
		for j := 0; j <= target; j++ {
			if nums[i-1] == j {
				dp[i][j] = 1
			}
		}
	}

	//dp
	for i := 1; i < len(nums)+1; i++ {
		for j := 0; j <= target; j++ {
			if j >= nums[i-1] {
				dp[i][j] += dp[i-1][j]
				if !(i-1 == 0 && j-nums[i-1] == 0) {
					dp[i][j] += dp[i-1][j-nums[i-1]]
				}
			} else {
				dp[i][j] = dp[i-1][j]
			}
		}
	}
	return dp[len(nums)][target]
}
