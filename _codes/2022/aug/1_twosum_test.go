package aug

import (
	"fmt"
	"testing"
)

func twoSum(nums []int, target int) []int {
	numsmap := make(map[int]struct{})
	for _, v := range nums {
		numsmap[v] = struct{}{}
	}
	for _, v := range nums {
		if _, ok := numsmap[target-v]; ok {
			return []int{target, target - v}
		}
	}
	return []int{-1, -1}

}

func TestSum(t *testing.T) {
	fmt.Println(twoSum([]int{1, 2, 4}, 3))
}
