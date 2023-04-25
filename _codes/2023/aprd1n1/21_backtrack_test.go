package main

import (
	"fmt"
	"testing"
)

func Test131(t *testing.T) {
	fmt.Println(partition("aabb"))
}

var ans [][]string

func partition(s string) [][]string {
	ans = [][]string{}

	left := s
	chosen := []string{}
	backtrack(left, chosen)
	return ans
}

func backtrack(left string, chosen []string) {
	if left == "" {
		tmp := make([]string, len(chosen))
		copy(tmp, chosen)
		ans = append(ans, tmp)
		return
	}

	for i := 0; i < len(left); i++ {
		cur := left[0 : i+1]
		if IsPalindrome(cur) {
			tmp := left
			chosen = append(chosen, cur)
			left = left[i+1:]
			backtrack(left, chosen)
			left = tmp //go back to the previous one
			chosen = chosen[:len(chosen)-1]
		}
	}
}

func IsPalindrome(s string) bool {
	left, end := 0, len(s)-1
	for left < end {
		if s[left] != s[end] {
			return false
		}
		left++
		end--
	}
	return true
}
