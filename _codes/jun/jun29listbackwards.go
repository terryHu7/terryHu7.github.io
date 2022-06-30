/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reversePrint(head *ListNode) []int {
	// if head == nil {
	// 	return nil
	// }

	return appendData(head)
}

func appendData(head *ListNode) []int {
	if head.Next != nil {
		list := appendData(head.Next)
		list = append(list, head.Val)
		return list
	}

	return []int{head.Val}
}

func reversePrint(head *ListNode) []int {
	if head == nil {
		return nil
	}

	var newHead *ListNode

	for head != nil {
		nexthead := head.Next
		head.Next = newHead
		newHead = head
		head = nexthead
	}
	res := []int{}
	for newHead != nil {
		res = append(res, newHead.Val)
		newHead = newHead.Next
	}

	return res
}

func reversePrint(head *ListNode) []int {
	if head == nil {
		return nil
	}

	res := []int{}
	for head != nil {
		res = append(res, head.Val)
		head = head.Next
	}

	for i, j := 0, len(res)-1; i < j; {
		res[i], res[j] = res[j], res[i]
		i++
		j--
	}

	return res
}

for i, j := 0, len(res)-1; i < j; {
	res[i], res[j] = res[j], res[i]
	i++
	j--
}