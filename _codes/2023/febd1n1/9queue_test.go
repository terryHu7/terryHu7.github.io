package main

type MyQueue struct {
	stack []int
	back  []int
}

/** Initialize your data structure here. */
func Constructor() MyQueue {
	return MyQueue{
		stack: make([]int, 0),
		back:  make([]int, 0),
	}
}

/** Push element x to the back of queue. */
func (q *MyQueue) Push(x int) {
	for len(q.back) != 0 {
		val := q.back[len(q.back)-1]
		q.back = q.back[:len(q.back)-1]
		q.stack = append(q.stack, val)
	}
	q.stack = append(q.stack, x)
}

/** Removes the element from in front of queue and returns that element. */
func (q *MyQueue) Pop() int {
	for len(q.stack) != 0 {
		val := q.stack[len(q.stack)-1]
		q.stack = q.stack[:len(q.stack)-1]
		q.back = append(q.back, val)
	}
	if len(q.back) == 0 {
		return 0
	}
	val := q.back[len(q.back)-1]
	q.back = q.back[:len(q.back)-1]
	return val
}

/** Get the front element. */
func (q *MyQueue) Peek() int {
	val := q.Pop()
	if val == 0 {
		return 0
	}
	q.back = append(q.back, val)
	return val
}

/** Returns whether the queue is empty. */
func (q *MyQueue) Empty() bool {
	return len(q.stack) == 0 && len(q.back) == 0
}
