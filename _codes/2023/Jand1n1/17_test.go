package main

import "testing"

func TestIntersection(t *testing.T) {

}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	l1, l2 := headA, headB
	for l1 != l2 {
		if l1 != nil {
			l1 = l1.Next
		} else {
			l1 = headB
		}
		if l2 != nil {
			l2 = l2.Next
		} else {
			l2 = headA
		}
	}
	return l1
}

// func detectCycle(head *ListNode) *ListNode {

// }

// d->x->x->x
// if you wanna delete the first one
// whether deleting it from the front or back, you always need the dummy head.

// should make it clear how to delete it from the front first

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{}
	dummy.Next = head

	fast, slow := dummy, dummy
	for i := n; i > 0; i-- {
		fast = fast.Next
	}
	for fast.Next != nil {
		fast = fast.Next
		slow = slow.Next
	}
	if slow.Next != nil && slow.Next.Next != nil {
		slow.Next = slow.Next.Next
	} else {
		slow.Next = nil
	}
	return dummy.Next
}
