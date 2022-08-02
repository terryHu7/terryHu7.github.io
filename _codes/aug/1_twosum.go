package main

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

func main() {

}
