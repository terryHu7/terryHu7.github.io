package main

import (
	"fmt"
	"testing"
)

func Test20(t *testing.T) {
	fmt.Println(isValid(")"))
}

func isValid(s string) bool {
	cmap := map[byte]byte{
		')': '(',
		'}': '{',
		']': '[',
	}
	stack := []byte{}

	for i, _ := range s {
		if s[i] == '(' || s[i] == '{' || s[i] == '[' {
			stack = append(stack, s[i])
		} else if len(stack) > 0 && stack[len(stack)-1] == cmap[s[i]] {
			stack = stack[:len(stack)-1]
		} else {
			return false
		}
	}
	return len(stack) == 0
}

func removeDuplicates(s string) string {
	stack := []byte{}
	for i := range s {
		stack = append(stack, s[i])
		if len(stack) >= 2 {
			if stack[len(stack)-1] == stack[len(stack)-2] {
				stack = stack[:len(stack)-2]
			}
		}
	}
	return string(stack)
}
