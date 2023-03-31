package main

import "fmt"

func main() {
	type node struct {
		data  int
		left  *node
		right *node
	}
	fmt.Println(twoSum([]int{2, 7, 11, 19}, 9))
}
func twoSum(nums []int, target int) []int {
	numsMap := make(map[int]int)
	
	for i := 0; i < len(nums); i++ {
		numsMap[i] = i
	}
	for _, v := range numsMap {
		minus := target - v
		if k, ok := numsMap[minus]; ok {
			return []int{k, v}
		}
	}
	return nil
}
