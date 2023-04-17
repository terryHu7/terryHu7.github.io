package main

import (
	"fmt"
	"sort"
	"testing"
)

func Test56(t *testing.T) {
	fmt.Println(partitionLabels("ababcbacadefegdehijhklij"))
}

func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[j][0] == intervals[i][0] {
			return intervals[j][1] > intervals[i][1]
		}
		return intervals[j][0] > intervals[i][0]
	})

	if len(intervals) == 0 {
		return [][]int{}
	}
	res := [][]int{intervals[0]}

	for i := 1; i < len(intervals); i++ {
		last := res[len(res)-1]
		cur := intervals[i]

		if cur[0] <= last[1] {
			if cur[1] >= last[1] {
				res[len(res)-1] = []int{last[0], cur[1]}
			} else {
				res[len(res)-1] = []int{last[0], last[1]}
			}

		} else {
			res = append(res, cur)
		}
	}
	return res
}

func partitionLabels(s string) []int {
	var res []int
	var marks [26]int
	size, left, right := len(s), 0, 0
	for i := 0; i < size; i++ {
		marks[s[i]-'a'] = i
	}
	fmt.Println(marks)
	for i := 0; i < size; i++ {
		right = max(right, marks[s[i]-'a'])
		if i == right {
			res = append(res, right-left+1)
			left = i + 1
		}
	}
	return res
}

func max(a, b int) int {
	if a < b {
		a = b
	}
	return a
}
