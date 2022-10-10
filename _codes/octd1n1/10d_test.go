package octd1n1

import (
	"fmt"
	"testing"
)

func TestMonoQ(t *testing.T) {
	fmt.Println(maxSlidingWindow([]int{1, 3, -1, -3, 5, 3, 6, 7}, 3))
}

func maxSlidingWindow(nums []int, k int) []int {
	monotonicQ, res := []int{}, []int{}

	for right := 0; right < len(nums); right++ {
		//push
		for len(monotonicQ) > 0 && monotonicQ[len(monotonicQ)-1] <= nums[right] {
			monotonicQ = monotonicQ[:len(monotonicQ)-1]
		}
		monotonicQ = append(monotonicQ, nums[right])

		if right >= k-1 {
			//res
			res = append(res, monotonicQ[0])
			//pop
			if nums[right-k+1] == monotonicQ[0] {
				monotonicQ = monotonicQ[1:]
			}
		}

	}
	return res
}
