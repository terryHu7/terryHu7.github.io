package aug

import (
	. "fmt"
	"testing"
)

func TestRangeCopy(t *testing.T) {
	a := []int{1, 2, 3}
	n := []int{}

	for _, v := range a {
		a = []int{}
		n = append(n, v)
	}
	Println(a, n) // [] [1 2 3]
}
