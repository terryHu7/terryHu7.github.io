package main

import (
	"fmt"
	"testing"
)

func Test22(t *testing.T) {
	fmt.Println(maxSubArray([]int{5, 4, -1, 7, 8}))
}

func findLengthOfLCIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	dp := make([]int, len(nums))
	dp[0] = 1
	res := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			dp[i] = dp[i-1] + 1
			if dp[i] > res { //update res
				res = dp[i]
			}
		} else {
			dp[i] = 1
		}
	}
	return res
}

//[-2 1 -3 4 -1 2 1 -5 4]
func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	dp := make([]int, len(nums))
	dp[0] = nums[0]

	for i := 1; i < len(nums); i++ {
		dp[i] = max(dp[i-1]+nums[i], nums[i]) //?
	}

	res := nums[0]
	return res
}
