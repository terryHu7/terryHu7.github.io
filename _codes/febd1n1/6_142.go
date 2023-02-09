package main

import "testing"

type ListNode struct {
	Val  int
	Next *ListNode
}

func Test142(t *testing.T) {

}

func detectCycle(head *ListNode) *ListNode {
	fast, slow := head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if slow == fast {
			for slow != head {
				slow = slow.Next
				head = head.Next
			}
			return slow
		}
	}
	return nil
}
