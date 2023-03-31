package octd1n1

import (
	"fmt"
	"testing"
)

var res bool
var used []bool
var bucket_target int

func TestKSubsets(t *testing.T) {
	fmt.Println(KSubsets([]int{4, 3, 2, 3, 5, 2, 1}, 4))
}

func KSubsets(nums []int, k int) bool {
	//init
	res, used, bucket_target = false, make([]bool, len(nums)), 0
	//exception
	if k > len(nums) {
		return false
	}
	var sum int
	for _, n := range nums {
		sum += n
	}
	if sum%k != 0 {
		return false
	}
	bucket_target = sum / k

	//
	backtrack(k, 0, nums, 0)
	return res
}

func backtrack(bucket_index int, bucket_sum int, nums []int, start int) {
	// only if
	if bucket_index == 0 {
		res = true
	}

	for i := start; i < len(nums); i++ {
		if used[i] {
			continue
		}
		if nums[i]+bucket_sum > bucket_target {
			continue
		}

		used[i] = true
		bucket_sum += nums[i]
		backtrack(bucket_index, bucket_sum, nums, i+1)
		used[i] = false
		bucket_sum -= nums[i]
	}

	if bucket_sum == bucket_target {
		backtrack(bucket_index-1, 0, nums, 0)
	}
}
