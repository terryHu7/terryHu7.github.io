package main

import "fmt"

func main() {
	ch := make(chan int)
	// go func() {
	// 	ch <- 1
	// }()
	ch <- 1
	val := <-ch
	fmt.Println(val)
}
