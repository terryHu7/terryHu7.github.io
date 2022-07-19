package main

import "fmt"

var Queue []int

func Enqueue(data int) {
	Queue = append(Queue, data)
}

func Dequeue() int {
	if len(Queue) == 0 {
		return -1
	} else {
		drop := Queue[0]
		Queue = Queue[1:] //???
		return drop
	}

}

func main() {
	Enqueue(0)
	Enqueue(1)
	Enqueue(2)
	fmt.Println(Queue)
	drop1 := Dequeue()
	drop2 := Dequeue()
	drop3 := Dequeue()
	fmt.Println(Queue, drop1, drop2, drop3)
}
