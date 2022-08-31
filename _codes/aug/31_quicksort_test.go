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

//***//
func partition(nums []int, start, end int) int {
	p := start
	for cur := start; cur < end; cur++ {
		if nums[cur] < nums[end] {
			nums[p], nums[cur] = nums[cur], nums[p]
			p++
		}
	}
	nums[p], nums[end] = nums[end], nums[p]
	return p
}

func TestQuickSort(t *testing.T) {
	fmt.Println(QuickSort([]int{3, 7, 2, 4, 9, 10, 6}, 0, 6))
}
