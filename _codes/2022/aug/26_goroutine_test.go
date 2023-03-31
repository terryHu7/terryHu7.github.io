package aug

import (
	"fmt"
	"testing"
	"time"
)

func TestGoroutine(t *testing.T) {
	// for i := 0; i < 10; i++ {
	// 	fmt.Println(i)
	// 	go fmt.Println(i)
	// }

	// time.Sleep(time.Second * time.Duration(1))

	c := make(chan string)
	s := ""
	go func() {
		s = <-c
	}()

	c <- "a"
	time.Sleep(time.Second)
	fmt.Println(s) // "a"
}

//git branch test
