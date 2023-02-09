package main

import (
	"fmt"
	"testing"
	"time"
)

var TChan chan int = make(chan int, 10)

var Queue []int

func Test7(t *testing.T) {

	is := false
	go func() {
		for {
			select {
			case t := <-TChan:
				Queue = append(Queue, t)

				if is == false {
					go func() {
						for len(Queue) != 0 {
							is = true

							time.Sleep(3 * time.Second)
							fmt.Println(Queue)

							is = false
							Queue = Queue[1:]
						}
					}()
				} else {
					fmt.Println("true")
				}

			}
		}
	}()

	TChan <- 1
	time.Sleep(time.Second)
	TChan <- 2
	time.Sleep(5 * time.Second)
}
