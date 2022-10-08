package sep_d1n1

import (
	"fmt"
	"testing"
)

func Test30D(t *testing.T) {
	var a int
	// if a, err := helper(); err == nil {
	// 	fmt.Println(a)
	// }
	a, _ = helper()
	fmt.Println(a)
}

func helper() (int, error) {
	return 100, nil
}
