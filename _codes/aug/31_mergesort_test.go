package aug

import (
	"fmt"
	"testing"
)

func MergeSort(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}
	mid := len(nums) / 2
	left := MergeSort(nums[:mid])
	right := MergeSort(nums[mid:])
	result := merge(left, right)
	return result
}
func merge(left, right []int) (result []int) {
	l := 0
	r := 0
	for l < len(left) && r < len(right) {
		if left[l] > right[r] {
			result = append(result, right[r])
			r++
		} else {
			result = append(result, left[l])
			l++
		}
	}
	result = append(result, left[l:]...)
	result = append(result, right[r:]...)
	return
}

func TestMergeSort(t *testing.T) {
	fmt.Println(MergeSort([]int{3, 7, 2, 4, 9, 10, 6}))
}
