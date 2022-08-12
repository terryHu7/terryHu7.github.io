package aug

import (
	"fmt"
	"strings"
	"testing"
)

func TestOffer50(t *testing.T) {
	res := firstUniqChar("dda")
	fmt.Println(string(res))
}

func firstUniqChar(s string) byte {
	sbytes := []byte(s)
	sexcept, before := []byte(""), []byte("")
	//strings.Contain()
	for i, w := range sbytes {
		if i == (len(sbytes) - 1) {
			sexcept = sbytes[:i]
		} else {
			copy(before, sbytes[:i]) //append，copy是真的坑
			sexcept = append(before, sbytes[i+1:]...)
		}
		if !strings.Contains(string(sexcept), string(w)) {
			return w
		}
	}
	return ' '
}
