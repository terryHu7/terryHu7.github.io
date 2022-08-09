package aug

import (
	"fmt"
	"strings"
	"testing"
)

func TestSplit(t *testing.T) {
	s := "It's all good man."
	sp := strings.Split(s, " ")
	fmt.Println(sp)
}
