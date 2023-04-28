package main

import (
	"fmt"
	"testing"
)

func TestChannel(t *testing.T) {

	tchan := make(chan int, 1)
	tchan <- 1
	// select {
	// default:
	// 	fmt.Println("2")
	// case <-tchan:
	// 	fmt.Println("1")
	// }
	<-tchan
	fmt.Println(len(tchan))
}
