package main

import "fmt"

type Node struct {
	Val  int
	Next *Node
}

type Stack struct {
	Top *Node
}

func (s *Stack) Push(data int) {
	newnode := &Node{data, nil}
	newnode.Next = s.Top
	s.Top = newnode
}

func (s *Stack) Pop() int {
	outnode := s.Top
	if outnode != nil {
		s.Top = outnode.Next
	} else {
		return -1
	}
	return outnode.Val
}

func main() {
	s := &Stack{}
	s.Push(0)
	s.Push(1)
	s.Push(2)
	out1 := s.Pop()
	out2 := s.Pop()
	out3 := s.Pop()
	out4 := s.Pop()
	fmt.Println(s, out1, out2, out3, out4)
}
