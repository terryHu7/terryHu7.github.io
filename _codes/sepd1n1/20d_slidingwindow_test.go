package sep_d1n1

import (
	"fmt"
	"math"
	"testing"
)

//			 WRONG!//////////////////////////////////////////
func TestSlidingWin(t *testing.T) {
	fmt.Println(minWindow("ADOBECODEBANC", "ABC"))
}

func minWindow(s string, t string) string {

	target, window := make(map[byte]int), make(map[byte]int)
	for i := 0; i < len(t); i++ {
		target[t[i]]++
	}

	left, right, oknum := 0, 0, 0
	res_start, res_len := 0, math.MaxInt32

	for right < len(s) { //right++
		r := s[right]
		right++
		// if _, ok := target[r]; ok {
		window[r]++
		if target[r] == window[r] {
			oknum++
		}
		// }

		for oknum == len(t) { //left++
			l := s[left]
			if window[r] == target[r] { //record res before left says byebye
				if right-left < res_len {
					res_start = left
					res_len = right - left
				}
				oknum--
			}
			window[l]-- //byebye
			left++
		}

	}

	if res_len == math.MaxInt32 {
		return ""
	} else {
		return s[res_start : res_start+res_len]
	}
}
