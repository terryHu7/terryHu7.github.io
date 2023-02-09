package main

import (
	"fmt"
	"testing"
)

var TestChan chan struct{} = make(chan struct{})

func TestTMP(t *testing.T) {
	go func() {
		TestChan <- struct{}{}
	}()

	select {
	case <-TestChan:
		fmt.Println("end")
	}

	fmt.Println("x")
}
