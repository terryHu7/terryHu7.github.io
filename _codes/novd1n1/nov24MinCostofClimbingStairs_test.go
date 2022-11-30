package main

import (
	"fmt"
	"testing"
)

func TestMinCostofClimbingStairs(t *testing.T) {
	fmt.Println(MinCostofClimbingStairs([]int{10, 15, 20}))
}

func MinCostofClimbingStairs(cost []int) int {
	//dp[i],i stands for the num of stairs, dp[i] stands for min cost
	//dp[0]=0, dp[1]=cost[0], dp[2]=min(dp[1]+cost[1],dp[0]+cost[0])

	// //1. climb from stair 0
	// if len(cost) == 0 {
	// 	return 0
	// }
	// if len(cost) == 1 {
	// 	return cost[0]
	// }
	// dp := make([]int, len(cost)+1)
	// dp[0], dp[1] = 0, cost[0]
	// for i := 2; i <= len(cost); i++ {
	// 	dp[i] = min(dp[i-1]+cost[i-1], dp[i-2]+cost[i-2])
	// }
	if len(cost) == 0 || len(cost) == 1 {
		return 0
	}
	dp := make([]int, len(cost)+1)
	dp[0], dp[1] = 0, 0
	for i := 2; i <= len(cost); i++ {
		dp[i] = min(dp[i-1]+cost[i-1], dp[i-2]+cost[i-2])
	}
	return dp[len(cost)]
}
func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
