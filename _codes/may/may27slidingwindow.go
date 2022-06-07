package main

import "fmt"

func main() {
	fmt.Println(lengthOfLongestSubstring("abcbanpf"))
}

func lengthOfLongestSubstring(s string) int {
	// if len(s) == 0 {
	// 	return 0
	// }
	win := make(map[byte]int)
	left, right := 0, 0
	ans := 1
	for right < len(s) {
		sr := s[right]
		right++
		win[sr]++
		//narrow down
		for win[sr] > 1 { //dont know where it is, only know there's a repeat
			sl := s[left]
			left++
			win[sl]--
		}
		ans = max(right-left, ans)
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
