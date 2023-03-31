package main

import (
	"fmt"
	"testing"
)

func TestStockProfit(t *testing.T) {
	fmt.Println(maxProfit([]int{7, 1, 5, 3, 6, 4}))
}

func maxProfit(prices []int) int {
	min, max :=
		func(a, b int) int {
			if a <= b {
				return a
			}
			return b
		},
		func(a, b int) int {
			if a >= b {
				return a
			}
			return b
		}

	in, res := prices[0], 0
	for i := 0; i < len(prices); i++ {
		in = min(in, prices[i])
		res = max(res, prices[i]-in)
	}
	return res
}
