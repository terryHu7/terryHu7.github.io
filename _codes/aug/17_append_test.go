package aug

import (
	"fmt"
	"testing"
)

func TestAppend(t *testing.T) {
	a := []int{1, 2, 3}
	b := []int{4, 5, 6}
	c := append(a[1:2], b...)
	fmt.Println(a, b, c)
}
