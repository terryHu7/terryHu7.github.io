package aug

import (
	"fmt"
	"testing"
)

func QuickSort(nums []int, start, end int) []int {
	if start < end {
		pivot := partition(nums, start, end)
		QuickSort(nums, 0, pivot-1)
		QuickSort(nums, pivot+1, end)
	}
	return nums
}

func partition(nums []int, start, end int) int {
	p := nums[end]
	i := start
	for j := start; j < end; j++ {
		if nums[j] < p {
			nums[i], nums[j] = nums[j], nums[i]
			i++
		}
	}
	nums[i], nums[end] = nums[end], nums[i]
	return i
}

func TestQuickSort(t *testing.T) {
	fmt.Println(QuickSort([]int{3, 7, 2, 4, 9, 10, 6}, 0, 6))
}
