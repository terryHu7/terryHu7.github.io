package main

import (
	"fmt"
	"sort"
	"testing"
)

func TestBiscuit(t *testing.T) {
	fmt.Println(largestSumAfterKNegations([]int{4, 2, 3}, 1))
}

func findContentChildren(g []int, s []int) int {
	children, biscuits := g, s

	sort.Ints(children)
	sort.Ints(biscuits)

	//children:[3,6], biscuit:[2,3,5]
	cur := 0
	for bindex := 0; bindex < len(biscuits); bindex++ {
		if cur == len(children) { //return asap.
			return cur
		}
		if biscuits[bindex] >= children[cur] { //*
			cur++
		}
	}
	return cur
}

/*
示例 1：
* 输入：A = [4,2,3], K = 1
* 输出：5
* 解释：选择索引 (1,) ，然后 A 变为 [4,-2,3]。

示例 2：
* 输入：A = [3,-1,0,2], K = 3
* 输出：6
* 解释：选择索引 (1, 2, 2) ，然后 A 变为 [3,1,0,2]。

示例 3：
* 输入：A = [2,-3,-1,5,-4], K = 2
* 输出：13
* 解释：选择索引 (1, 4) ，然后 A 变为 [2,3,-1,5,4]。
*/

func largestSumAfterKNegations(nums []int, K int) int {
	findminIndex := func(nums []int) int {
		mindex := 0
		for i, n := range nums {
			if n < nums[mindex] {
				mindex = i
			}
		}
		return mindex
	}

	for i := 0; i < K; i++ {
		minindex := findminIndex(nums)
		nums[minindex] = -nums[minindex]
	}

	sum := 0
	for _, n := range nums {
		sum += n
	}
	return sum
}

func lemonadeChange(bills []int) bool {
	five, ten, twenty := 0, 0, 0
	for _, b := range bills {
		switch b {
		case 5:
			five++
		case 10:
			ten++
			if five > 0 {
				five--
			} else {
				return false
			}
		case 20:
			if ten > 0 && five > 0 {
				ten--
				five--
				twenty++
			} else if five >= 3 {
				five -= 3
			} else {
				return false
			}
		}
	}
	return true
}
