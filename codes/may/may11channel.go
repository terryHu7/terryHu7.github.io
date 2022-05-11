package main

import (
	"fmt"
	"time"
)

func main() {
	p := fmt.Println
	c := make(chan int)
	go f1(c)
	p("<-c: ", <-c)
	p("<-c: ", <-c)
	time.Sleep(1 * time.Second)
	p("DONE!")
}

func f1(c chan int) {
	p := fmt.Println
	p("into go1")
	defer p("out go1")

	for i := 0; i < 3; i++ {
		p(i)
	}
	c <- 1
	p("c<- 1")
	c <- 2
	p("c<- 2")
}
