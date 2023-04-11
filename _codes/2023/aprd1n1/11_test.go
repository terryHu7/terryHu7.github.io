package main

import (
	"fmt"
	"sort"
	"testing"
)

func TestWiggle(t *testing.T) {
	fmt.Println(candy([]int{1, 2, 3, 3, 3, 2, 1}))
}

func wiggleMaxLength(nums []int) int {
	// len=0/1/2
	l := len(nums)
	if l <= 1 {
		return l
	}
	res := 0
	prevDiff := nums[1] - nums[0]
	if prevDiff != 0 {
		res = 2
	} else {
		res = 1
	}
	if l == 2 {
		return res
	}

	// len>=3
	for i := 2; i < len(nums); i++ {
		diff := nums[i] - nums[i-1]
		if diff > 0 && prevDiff <= 0 || diff < 0 && prevDiff >= 0 {
			res++
			prevDiff = diff
		}
	}
	return res
}

func monotoneIncreasingDigits(N int) int {
	d, r := N, 0
	Nums := []int{}
	for d != 0 {
		r = d % 10
		Nums = prepend(Nums, r)
		d = d / 10
	}

	for i := len(Nums) - 1; i > 0; i-- {
		if Nums[i-1] > Nums[i] {
			Nums[i-1]--
			for j := i; j < len(Nums); j++ {
				Nums[j] = 9
			}
		}
	}
	res := 0
	for _, n := range Nums {
		res = res*10 + n
	}
	return res
}

func prepend(slice []int, element int) []int {
	newSlice := make([]int, len(slice)+1)
	newSlice[0] = element
	copy(newSlice[1:], slice)
	return newSlice
}

//?
func candy(ratings []int) int {
	candies := []int{}
	for i := 0; i < len(ratings); i++ {
		candies = append(candies, 1)
	}

	for i := 0; i < len(ratings)-1; i++ {
		if ratings[i+1] > ratings[i] {
			candies[i+1]++
		}
	}
	fmt.Println(candies)
	for i := len(ratings) - 1; i > 0; i-- {
		if ratings[i-1] > ratings[i] {
			if candies[i-1] < candies[i]+1 {
				candies[i-1] = candies[i] + 1
			}
		}
	}

	sum := 0
	for _, c := range candies {
		sum += c
	}
	fmt.Println(candies)
	return sum
}

func reconstructQueue(people [][]int) [][]int {
	//先将身高从大到小排序，确定最大个子的相对位置
	sort.Slice(people, func(i, j int) bool {
		if people[i][0] == people[j][0] {
			return people[i][1] < people[j][1] //这个才是当身高相同时，将K按照从小到大排序
		}
		return people[i][0] > people[j][0] //这个只是确保身高按照由大到小的顺序来排，并不确定K是按照从小到大排序的
	})
	//再按照K进行插入排序，优先插入K小的
	result := make([][]int, 0)
	for _, info := range people {
		result = append(result, info)
		copy(result[info[1]+1:], result[info[1]:]) //将插入位置之后的元素后移动一位（意思是腾出空间）
		result[info[1]] = info                     //将插入元素位置插入元素
	}
	return result
}
