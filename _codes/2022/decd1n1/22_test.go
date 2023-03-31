package main

import (
	"fmt"
	"testing"
)

func Test22(t *testing.T) {
	fmt.Println(maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
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
	res := nums[0]
	for i := 1; i < len(nums); i++ {
		dp[i] = max(dp[i-1]+nums[i], nums[i])
		if dp[i] > res {
			res = dp[i]
		}
	}
	fmt.Println(dp)
	return res
}

// func max3(elements ...int) int {
// 	res := elements[0]
// 	for _, e := range elements {
// 		if e > res {
// 			res = e
// 		}
// 	}
// 	return res
// }
