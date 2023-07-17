package main

import "testing"

func TestTemove(t *testing.T) {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{
		Val:  0,
		Next: head,
	}
	start := head
	for i := n; i > 0; i-- {
		start = start.Next
	}

	for start != nil && start.Next != nil {
		head = head.Next
		start = start.Next
	}

	if head.Next != nil {
		head = head.Next.Next
	}
	return dummy.Next
}
