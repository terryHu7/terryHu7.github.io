package main

import (
	"fmt"
	"testing"
)

//nums = [0,1,2,2,3,0,4,2], val = 2,

func TestSwap(t *testing.T) {
	fmt.Println(removeElement([]int{0, 1, 2, 2, 3, 0, 4, 2}, 2))
}

func removeElement(nums []int, val int) int {
	slow, fast := 0, 0
	for ; fast < len(nums); fast++ {
		if nums[fast] != val {
			nums[slow], nums[fast] = nums[fast], nums[slow]
			slow++
		}
	}
	fmt.Println(nums)
	return slow
}
