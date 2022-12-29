package tools

import (
	"fmt"
	"time"
)

func Ticker() {
	ticker := time.NewTicker(time.Second)
	done := make(chan bool)

	go func() {
		for {
			select {
			case <-ticker.C: //block
				{
					fmt.Println("tik tik...")
				}
			case <-done:
				{
					return
				}
			}
		}
	}()
	time.Sleep(1 * time.Second)
	ticker.Stop()
	done <- true
	fmt.Println("tick finished")
}
