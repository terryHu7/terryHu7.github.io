package aug

import (
	"fmt"
	"testing"
)

func TestReverse(t *testing.T) {
	s := []rune("background")
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	// i, j := 0, len(s)-1
	// for i < j {
	// 	s[i], s[j] = s[j], s[i]
	// 	i++
	// 	j--
	// }
	fmt.Println(string(s))
}
