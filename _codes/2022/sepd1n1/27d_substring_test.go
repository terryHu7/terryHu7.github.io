package sep_d1n1

import (
	"fmt"
	"testing"
)

func TestSubString(t *testing.T) {
	fmt.Println(slidingWindow("ab", "eidboaoo"))
}

func slidingWindow(s1, s2 string) bool {
	left, right := 0, 0
	okkey := 0
	res := false
	target, window := make(map[byte]int), make(map[byte]int)

	for i := 0; i < len(s1); i++ {
		target[s1[i]]++
	}

	for right < len(s2) {
		r := s2[right]
		right++
		window[r]++
		if window[r] == target[r] {
			okkey++
		}
		///***///
		if right-left == len(s1) { //len-1?
			if okkey == len(target) {
				res = true
			}
			///***///
			l := s2[left]
			left++
			if window[l] == target[l] {
				okkey--
			}
			window[l]--
		}
	}
	return res
}
