package main

import "fmt"

type Node struct {
	Val  int
	Next *Node //invalid recursive type
}

type Queue struct {
	Head *Node
	Tail *Node
	// if we dont have Tail,
	// everytime we want to enqueue, we have to traverse from beginning
}

func (q *Queue) Enqueue(data int) {

	newnode := &Node{data, nil}

	if q.Head == q.Tail && q.Head == nil { //1st, nothing
		q.Head = newnode
		q.Tail = newnode
	} else if q.Tail != nil { //2nd, 3rd...
		q.Tail.Next = newnode
		q.Tail = newnode
	}
	// q.Tail.Next.Val = data
	// Since q.Tail is nil, nil.Next leads to "nil pointer dereference"
}

func (q *Queue) Dequeue() int {
	dropnode := q.Head

	if q.Head == q.Tail && q.Head != nil { //last time
		q.Head, q.Tail = nil, nil
		return dropnode.Val

	} else if q.Head != nil {
		q.Head = q.Head.Next
		return dropnode.Val
	} else {
		return -1
	}

}

func main() {
	q := &Queue{}
	q.Enqueue(0)
	q.Enqueue(1)
	q.Enqueue(2)
	drop1 := q.Dequeue()
	drop2 := q.Dequeue()
	drop3 := q.Dequeue()
	fmt.Println(q, drop1, drop2, drop3)
}
