package main

import (
	"fmt"
	"testing"
)

var A int = 1

func TestClosure(t *testing.T) {
	// tmp()
	// fmt.Println(A)
	res := 0
	f := func(x int) int {
		res = (res + x) * 2
		return res
	}

	for i := 0; i < 10; i++ {
		fmt.Println(f(i))
	}
}
func tmp() {
	A = 2
	fmt.Println(A)
}
