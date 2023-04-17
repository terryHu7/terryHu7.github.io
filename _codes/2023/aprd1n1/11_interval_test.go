package main

import "testing"

func TestInterval(t *testing.T) {

}

func canJump(nums []int) bool {
	cover := 0
	end := len(nums) - 1

	for i := 0; i <= cover; i++ {
		if nums[i]+i > cover {
			cover = nums[i] + i
		}
		if cover >= end {
			return true
		}
	}
	return false
}
