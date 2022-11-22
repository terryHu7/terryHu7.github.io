package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	wg.Add(1)
	for i := 0; i < 5; i++ {
		go func(wgptr *sync.WaitGroup) {
			wg.Wait()
			fmt.Println("resume")
		}(&wg)
	}

	time.Sleep(3 * time.Second)
	wg.Done()
	time.Sleep(1 * time.Second)

	v := "abc"
	switch v {
	default:
		{
			fmt.Println("default")
		}
	case "abc":
		{
			fmt.Println("abc")
		}
	}

}
