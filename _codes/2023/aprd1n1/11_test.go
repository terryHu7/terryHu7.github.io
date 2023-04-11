package main

import (
	"fmt"
	"sort"
	"testing"
)

func TestWiggle(t *testing.T) {
	example := make([]int, 5)
	example[0] = 1
	example[1] = 2
	example[2] = 3
	example[3] = 4
	example[4] = 5
	fmt.Println(example)
	copy(example[1:], example[0:])
	fmt.Println(example)

	// fmt.Println(candy([]int{1, 2, 3, 3, 3, 2, 1}))
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
	candies := make([]int, len(ratings))
	for i := 0; i < len(candies); i++ {
		candies[i] = 1
	}

	for i := 0; i < len(ratings)-1; i++ {
		if ratings[i+1] > ratings[i] {
			candies[i+1] = candies[i] + 1
		}
	}
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
	// fmt.Println(candies)
	return sum
}

func reconstructQueue(people [][]int) [][]int {
	sort.Slice(people, func(i, j int) bool {
		if people[j][0] == people[i][0] { // j is the former element
			return people[j][1] > people[i][1]
		}
		return people[j][0] < people[i][0] // "true" means "yes, switch"
	})

	result := make([][]int, 0)

	//For example:
	//[7,0],[6,1],[7,1]

	//insert:[5,0]
	for _, p := range people {
		result = append(result, p)           //[7,0],[6,1],[7,1],[5,0]
		copy(result[p[1]+1:], result[p[1]:]) //[7,0],[7,0],[6,1],[7,1]
		result[p[1]] = p                     //[5,0],[7,0],[6,1],[7,1]
	}

	return result
}
