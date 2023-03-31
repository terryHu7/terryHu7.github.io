package main

import "fmt"

func main() {
	head := newList([]int{1, 1, 1, 1, 1})
	fmt.Println(hasCycle(head)) // false
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func newList(nums []int) *ListNode {
	n := len(nums)
	if n == 0 {
		return nil
	}
	head := &ListNode{Val: nums[0], Next: nil}
	cur := head
	for i := 1; i < n; i++ {
		newNode := &ListNode{Val: nums[i], Next: nil}
		cur.Next = newNode
		cur = newNode
	}
	return head
}

//pointer is an address.
// var map141 = map[*ListNode]int{}

// func hasCycle(head *ListNode) bool {
// 	if head == nil {
// 		return false
// 	} else {
// 		map141[head] = 0
// 	}

// 	cur := head
// 	for {
// 		cur = cur.Next //:=
// 		if cur == nil {
// 			return false
// 		} else {
// 			if _, exist := map141[cur]; exist {
// 				return true
// 			} else {
// 				map141[cur] = 0
// 			}
// 		}
// 	}

// }

var map141 = map[string]string{}

func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	} else {
		map141[fmt.Sprint(head)] = fmt.Sprint(head)
	}

	cur := head
	for {
		cur = cur.Next //:=
		if cur == nil {
			return false
		} else {
			scur := fmt.Sprint(cur)
			if _, exist := map141[scur]; exist {
				return true
			} else {
				map141[scur] = scur
			}
		}
	}
}
