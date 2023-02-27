package main

import (
	"fmt"
	"strconv"
	"testing"
)

func Test150(t *testing.T) {
	tmp := []int{1, 2, 3}
	a := pop(&tmp)
	fmt.Println(tmp, a)
}

func evalRPN(tokens []string) int {
	stack := []int{}
	for _, c := range tokens {
		if val, err := strconv.Atoi(c); err == nil {
			stack = append(stack, val)
		} else {
			p1, p2, tmp := pop(&stack), pop(&stack), 0
			switch c {
			case "+":
				tmp = p2 + p1
			case "-":
				tmp = p2 - p1
			case "*":
				tmp = p2 * p1
			case "/":
				tmp = p2 / p1

			}
			stack = append(stack, tmp)
		}
	}
	return pop(&stack)
}

func pop(stackPtr *[]int) int {
	p := (*stackPtr)[len(*stackPtr)-1]
	*stackPtr = (*stackPtr)[:len(*stackPtr)-1]
	return p
}
