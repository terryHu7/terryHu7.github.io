package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	say2("hello", &wg)
	say2("world", &wg)
	// wg.Wait()//???
	fmt.Println("over!")
}

func say2(s string, waitGroup *sync.WaitGroup) {
	defer waitGroup.Done()

	for i := 0; i < 10; i++ {
		fmt.Println(s)
	}
}
