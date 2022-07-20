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
		Queue = Queue[1:] //???  []int{1}[1:] & []int{1}[1]
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

	test := []int{1}
	fmt.Println(test[1:])

	test2 := []int{}
	var test3 []int
	fmt.Println(test2 == nil)
	fmt.Printf("test2: %p, test3: %p", test2, test3)
}
