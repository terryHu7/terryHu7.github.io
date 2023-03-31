package main

import "fmt"

type Node struct {
	Val  int
	Next *Node
}

var Stack []int

func Push(data int) {
	Stack = append(Stack, data)
}

func Pop() int {
	if len(Stack) == 0 {
		return -1
	} else {
		lifo := len(Stack) - 1
		out := Stack[lifo]
		Stack = Stack[:lifo] //[:0]
		return out
	}

}

func main() {

	Push(0)
	Push(1)
	Push(2)
	out1 := Pop()
	out2 := Pop()
	out3 := Pop()
	out4 := Pop()
	fmt.Println(Stack, out1, out2, out3, out4)
}
