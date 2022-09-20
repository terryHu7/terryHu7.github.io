package sep_d1n1

import (
	"math"
	"testing"
)

func TestSlidingWin(t *testing.T) {

}

func minWindow(s string, t string) string {
	need, window := map[byte]int{}, map[byte]int{}
	for i := 0; i < len(t); i++ {
		need[t[i]]++
	}

	left, right := 0, 0
	valid := 0
	start, temp_len := 0, math.MaxInt32
	for right < len(s) {
		c := s[right]
		right++
		if need[c] != 0 {
			window[c]++
			if window[c] == need[c] {
				valid++
			}
		}

		for valid == len(need) {
			if right-left < temp_len {
				start = left
				temp_len = right - left
			}
			d := s[left]
			left++
			if need[d] != 0 {
				if window[d] == need[d] {
					valid--
				}
				window[d]--
			}
		}
	}
	if temp_len == math.MaxInt32 {
		return ""
	} else {
		return s[start : start+temp_len]
	}
}
