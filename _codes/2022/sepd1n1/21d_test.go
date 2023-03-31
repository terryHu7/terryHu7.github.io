package sep_d1n1

import (
	"fmt"
	"testing"
)

func TestMap(t *testing.T) {
	nums := []int{1}
	nums_map := make(map[int]int)
	for _, n := range nums {
		nums_map[n]++
	}
	// for _, n := range nums {
	// 	if _, ok := nums_map[n]; ok {
	// 		nums_map[n]++
	// 	} else {
	// 		nums_map[n] = 1
	// 	}
	// }
	fmt.Printf("%+v\n", nums_map)
}
