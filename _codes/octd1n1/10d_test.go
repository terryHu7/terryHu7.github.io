package octd1n1

import (
	"fmt"
	"testing"
)

func TestMonoQ(t *testing.T) {
	fmt.Println(maxSlidingWindow([]int{1, 3, -1, -3, 5, 3, 6, 7}, 3))
}

func maxSlidingWindow(nums []int, k int) []int {
	monotonic_q := []int{}
	var res []int
	for right := 0; right < len(nums); right++ {

		for len(monotonic_q) > 0 && nums[right] >= monotonic_q[len(monotonic_q)-1] {
			monotonic_q = monotonic_q[:len(monotonic_q)-1]
		}
		monotonic_q = append(monotonic_q, nums[right])

		// fmt.Println(monotonic_q)

		if right >= k-1 {
			res = append(res, monotonic_q[0])
			if nums[right-k+1] == monotonic_q[0] {
				monotonic_q = monotonic_q[1:]
			}
		}

	}
	return res
}
