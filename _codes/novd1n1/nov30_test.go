package main

import (
	"fmt"
	"testing"
)

func TestLIS(t *testing.T) {
	fmt.Println(LIS([]int{10, 9, 2, 5, 3, 7, 101, 18}))
}

func LIS(nums []int) int {
	//[10,9,2,5,3,7,101,18]

	//Initialize
	dp := make([]int, len(nums))
	for i := 0; i < len(dp); i++ {
		dp[i] = 1
	}
	res := 1

	for i := 1; i < len(dp); i++ {
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
		//update res
		if dp[i] > res {
			res = dp[i]
		}
	}
	return res

}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
