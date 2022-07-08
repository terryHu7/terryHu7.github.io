package main

import (
	"fmt"
	"time"
)

type StructWithChan struct {
	callback *(chan interface{})
}

func main() {
	s1 := &StructWithChan{}
	fmt.Println(s1.callback == nil)

	s2 := &StructWithChan{
		callback: new(chan interface{}),
	}
	fmt.Println(s2.callback == nil)

	s3callback := make(chan interface{})
	s3 := &StructWithChan{
		callback: &s3callback,
	}
	ticker := time.NewTimer(time.Second * 3)
	// go func() {
	// 	time.Sleep(time.Second)
	// 	s3callback <- "hi"
	// }()

LOOP:
	for {
		select {
		case <-s3callback:
			{
				fmt.Println("hi")
				break LOOP
			}
		case <-ticker.C:
			{

				fmt.Println("3s finished, no answer")
				break LOOP
			}
		}
	}
	fmt.Println(*s3.callback, s3callback)
}
